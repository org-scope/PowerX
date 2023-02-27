package department

import (
	"PowerX/internal/svc"
	"PowerX/internal/types"
	"PowerX/internal/types/errorx"
	"PowerX/internal/uc"
	"PowerX/pkg/mapx"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetDepartmentTreeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDepartmentTreeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDepartmentTreeLogic {
	return &GetDepartmentTreeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDepartmentTreeLogic) GetDepartmentTree(req *types.GetDepartmentTreeRequest) (resp *types.GetDepartmentTreeReply, err error) {
	var userIds []int64
	depPage := l.svcCtx.UC.Department.FindManyDepartments(l.ctx, &uc.FindManyDepartmentsOption{
		RootId: req.DepId,
	})

	if depPage.Total == 0 {
		return nil, errorx.WithCause(errorx.ErrBadRequest, "未找到根部门")
	}

	for _, department := range depPage.List {
		for _, id := range department.LeaderIds {
			userIds = append(userIds, id)
		}
	}

	userPage := l.svcCtx.UC.Employee.FindManyEmployees(l.ctx, &uc.FindEmployeeOption{
		Ids: userIds,
	})
	userMap := mapx.MapByFunc(userPage.List, func(item *uc.Employee) (int64, *uc.Employee) {
		return item.ID, item
	})

	// make node
	var rootNode *types.DepartmentNode
	var voSlice []types.DepartmentNode
	voGroupByPid := make(map[int64][]types.DepartmentNode)
	for _, department := range depPage.List {
		leaders := make([]types.SimpleEmployee, 0)
		for _, id := range department.LeaderIds {
			if u, ok := userMap[id]; ok {
				leaders = append(leaders, types.SimpleEmployee{
					Id:   u.ID,
					Name: u.Name,
				})
			}
		}
		node := types.DepartmentNode{
			Id:      department.ID,
			DepName: department.Name,
			Leaders: leaders,
		}
		voSlice = append(voSlice, node)
		if node.Id == req.DepId {
			rootNode = &node
		}
		if voGroupByPid[department.PId] == nil {
			voGroupByPid[department.PId] = make([]types.DepartmentNode, 0, 1)
		}
		voGroupByPid[department.PId] = append(voGroupByPid[department.PId], node)
	}

	if rootNode == nil {
		return nil, errorx.WithCause(errorx.ErrBadRequest, "未找到根部门")
	}

	var makeChildren func(currentNode types.DepartmentNode) []types.DepartmentNode
	makeChildren = func(currentNode types.DepartmentNode) []types.DepartmentNode {
		if nodes, ok := voGroupByPid[currentNode.Id]; ok {
			for i := range nodes {
				nodes[i].Children = makeChildren(nodes[i])
			}
			return nodes
		} else {
			return make([]types.DepartmentNode, 0)
		}
	}
	rootNode.Children = makeChildren(*rootNode)
	resp = &types.GetDepartmentTreeReply{
		DepTree: *rootNode,
	}
	return
}