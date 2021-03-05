package ldb

type Method string

const (
	List   Method = "list"
	Get    Method = "get"
	Put    Method = "put"
	Delete Method = "delete"
)
