// Package service
package service

type Service interface{
	Active()
	Shutdown()
}

