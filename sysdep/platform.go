package sysdep

import "cbc/models"

type Platform interface {
	GetTypeTable() models.TypeTable
}
