package model

import (
	"time"
)

type TAlbum struct {
	ID           int        `json:"id"`                                    // id
	UserId       int        `json:"user_id"`                               // user_id
	OrderNum     int        `json:"order_num"`                             // order_num
	Name         string     `json:"name"`                                  // name
	CreatedAt    time.Time  `json:"create_date" gorm:"column:create_date"` // create_date
	CreateUserID int        `json:"create_user_id"`                        // create_user_id
	UpdatedAt    time.Time  `json:"update_date" gorm:"column:update_date"` // update_date
	UpdateUserID int        `json:"update_user_id"`                        // update_user_id
	DeletedAt    *time.Time `json:"delete_date" gorm:"column:delete_date"` // delete_date
	DeleteUserID int        `json:"delete_user_id"`                        // delete_user_id
	DeleteFlg    bool       `json:"delete_flg"`                            // delete_flg
}
