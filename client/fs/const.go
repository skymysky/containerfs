// Copyright 2018 The Containerfs Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License.

package fs

import (
	"syscall"
	"time"

	"github.com/tiglabs/containerfs/fuse"

	"github.com/tiglabs/containerfs/proto"
)

const (
	RootInode = proto.RootIno
)

const (
	DIR_NLINK_DEFAULT     = 2
	REGULAR_NLINK_DEFAULT = 1
)

const (
	DefaultBlksize    = uint32(1) << 12
	DefaultMaxNameLen = uint32(256)
)

const (
	LookupValidDuration = 0
	AttrValidDuration   = 30 * time.Second
)

const (
	DefaultInodeExpiration = 120 * time.Second
	MaxInodeCache          = 10000000
)

const (
	DentryValidDuration = 5 * time.Second
)

const (
	DeleteExtentsTimeout = 600 * time.Second
)

func ParseError(err error) fuse.Errno {
	switch v := err.(type) {
	case syscall.Errno:
		return fuse.Errno(v)
	case fuse.Errno:
		return v
	default:
		return fuse.ENOSYS
	}
}

func ParseMode(mode uint32) fuse.DirentType {
	if proto.IsDir(mode) {
		return fuse.DT_Dir
	} else if proto.IsSymlink(mode) {
		return fuse.DT_Link
	}
	return fuse.DT_File
}
