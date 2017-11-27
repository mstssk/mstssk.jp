#!/bin/sh -eux

cd `dirname $0`

gb gae serve src
