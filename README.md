# go-append-annotation
go helper to add annotions in go sourcecode

## Problem
On write a Kubernetes Operator for a third party Application written in go, the use of their classes to create a CRD was not possible because missing kbuilder annotions.

## Solution
* Use git submodule to get go module local
* Replace submodule with local in go.mod
* Generate a config.yml define the annotations
* Run this program with "-config config.yml" to add annotions
