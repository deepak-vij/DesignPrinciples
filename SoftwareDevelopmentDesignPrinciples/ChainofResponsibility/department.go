// This is the handler interface

package main

type Department interface {
    execute(*Patient)
    setNext(Department)
}