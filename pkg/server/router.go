package server

type Method string

const GET Method = "GET"
const POST Method = "POST"
const PUT Method = "PUT"
const PATCH Method = "PATCH"
const DELETE Method = "DELETE"
const HEAD Method = "HEAD"
const CONNECT Method = "CONNECT"
const OPTIONS Method = "OPTIONS"
const TRACE Method = "TRACE"

type Router interface {
	Add(method Method)
}

type RouterImpl struct {
}