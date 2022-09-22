package main

type Entity interface {
	worldToLocal(Vec3) Vec3
}
