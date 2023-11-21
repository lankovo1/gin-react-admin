package system

import (
	"server/global"
	commonRequest "server/models/common/request"
	"server/models/system"
	"server/utils"
)

type RoleServices struct {
}

func (RoleService *RoleServices) GetRoleList(search commonRequest.Search[system.SysRole]) (roles []system.SysRole, total int64, err error) {
	limit, offset := utils.PageQuery(search.PageInfo)
	role := search.Condition
	tx := global.GRA_DB.Model(system.SysRole{}).Scopes(
		utils.SearchWhere("id", role.Id, false),
		utils.SearchWhere("role_name", role.RoleName, true))
	err = tx.Count(&total).Error
	if err != nil {
		return roles, 0, err
	}
	err = tx.Offset(offset).Limit(limit).Select("id,role_name,default_router_id").Order("id").Find(&roles).Error
	return roles, total, err
}

func (RoleService *RoleServices) UpdateRole(role system.SysRole) error {
	return global.GRA_DB.Where("id = ?", role.Id).Updates(&role).Error
}

func (RoleService *RoleServices) DeleteRole(id []uint) error {
	return global.GRA_DB.Delete(system.SysRole{}, &id).Error
}

func (RoleService *RoleServices) FindRoleById(role system.SysRole) (system.SysRole, error) {
	err := global.GRA_DB.Where("id = ?", role.Id).Select("id,role_name").Find(&role).Error
	return role, err
}
func (RoleService *RoleServices) InsertRole(role system.SysRole) error {
	return global.GRA_DB.Select("role_name").Create(&role).Error
}
func (RoleService *RoleServices) GetRoleMenuTree(id string) (routers []system.SysRouter, role system.SysRole, err error) {
	err = global.GRA_DB.Where("id = ?", id).Select("allow_router_id,default_router_id").Find(&role).Error
	if err != nil {
		return nil, role, err
	}
	err = global.GRA_DB.Order("id,router_order").Find(&routers).Error
	return routers, role, err
}
func (RoleService *RoleServices) GetRoleAuthority(id string) (role system.SysRole, err error) {
	err = global.GRA_DB.Where("id = ?", id).Select("allow_api_id").Find(&role).Error
	if err != nil {
		return role, err
	}
	return role, err
}
