// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// VmDao is the data access object for table vm.
type VmDao struct {
	table   string    // table is the underlying table name of the DAO.
	group   string    // group is the database configuration group name of current DAO.
	columns VmColumns // columns contains all the column names of Table for convenient usage.
}

// VmColumns defines and stores column names for table vm.
type VmColumns struct {
	Id          string //
	Uuid        string // 主机uuid
	Name        string //
	Hostname    string //
	OsType      string // 操作系统类型
	OsInfo      string // 操作系统信息
	Hosttype    string // 主机类型
	Networktype string // 网络类型
	PrivateIp   string //
	PublicIp    string //
	Created     string //
	Updated     string //
	Creater     string // 创建人
	Address     string //
	PeerId      string //
}

//  vmColumns holds the columns for table vm.
var vmColumns = VmColumns{
	Id:          "id",
	Uuid:        "uuid",
	Name:        "name",
	Hostname:    "hostname",
	OsType:      "os_type",
	OsInfo:      "os_info",
	Hosttype:    "hosttype",
	Networktype: "networktype",
	PrivateIp:   "private_ip",
	PublicIp:    "public_ip",
	Created:     "created",
	Updated:     "updated",
	Creater:     "creater",
	Address:     "address",
	PeerId:      "peer_id",
}

// NewVmDao creates and returns a new DAO object for table data access.
func NewVmDao() *VmDao {
	return &VmDao{
		group:   "default",
		table:   "vm",
		columns: vmColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *VmDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *VmDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *VmDao) Columns() VmColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *VmDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *VmDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *VmDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
