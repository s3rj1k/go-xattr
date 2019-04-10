package xattr // import "github.com/s3rj1k/go-xattr"

import (
	"os"

	"golang.org/x/sys/unix"
)

// Todo: Test on XFS
const (
	// FsIocGetFlags is FS_IOC_GETFLAGS from /usr/include/linux/fs.h (man ioctl_list)
	FsIocGetFlags = 0x80086601
	// FsIocSetFlags is FS_IOC_SETFLAGS from /usr/include/linux/fs.h (man ioctl_list)
	FsIocSetFlags = 0x40086602

	// FsImmutableFlag is FS_IMMUTABLE_FL from /usr/include/linux/fs.h
	FsImmutableFlag = 0x00000010 /* Immutable file */
	// FsAppendFlag is FS_APPEND_FL from /usr/include/linux/fs.h
	FsAppendFlag = 0x00000020 /* writes to file may only append */
	// FsNoDumpFlag is FS_NODUMP_FL from /usr/include/linux/fs.h
	FsNoDumpFlag = 0x00000040 /* do not dump file */
)

// GetAttr retrieves the attributes of a file on a linux filesystem.
func GetAttr(fd *os.File) (int, error) {
	return unix.IoctlGetInt(int(fd.Fd()), FsIocGetFlags)
}

// GetAttrFdPtr retrieves the attributes of a file on a linux filesystem, using file descriptor as uintptr directly.
func GetAttrFdPtr(fd uintptr) (int, error) {
	return unix.IoctlGetInt(int(fd), FsIocGetFlags)
}

// GetAttrFd retrieves the attributes of a file on a linux filesystem, using file descriptor as int directly.
func GetAttrFd(fd int) (int, error) {
	return unix.IoctlGetInt(fd, FsIocGetFlags)
}

// SetAttr sets the attributes of a file on a linux filesystem to the given value.
func SetAttr(fd *os.File, attr int) error {
	return unix.IoctlSetPointerInt(int(fd.Fd()), FsIocSetFlags, attr)
}

// SetAttrFdPtr sets the attributes of a file on a linux filesystem to the given value, using file descriptor as uintptr directly.
func SetAttrFdPtr(fd uintptr, attr int) error {
	return unix.IoctlSetPointerInt(int(fd), FsIocSetFlags, attr)
}

// SetAttrFd sets the attributes of a file on a linux filesystem to the given value, using file descriptor as uintptr directly.
func SetAttrFd(fd int, attr int) error {
	return unix.IoctlSetPointerInt(fd, FsIocSetFlags, attr)
}

// IsImmutable checks if file has immutable flag set.
func IsImmutable(fd *os.File) bool {
	if xattr, err := GetAttr(fd); err == nil && xattr&FsImmutableFlag != 0 {
		return true
	}

	return false
}

// IsImmutableFdPtr checks if file has immutable flag set, using file descriptor as uintptr directly.
func IsImmutableFdPtr(fd uintptr) bool {
	if xattr, err := GetAttrFdPtr(fd); err == nil && xattr&FsImmutableFlag != 0 {
		return true
	}

	return false
}

// IsImmutableFd checks if file has immutable flag set, using file descriptor as int directly.
func IsImmutableFd(fd int) bool {
	if xattr, err := GetAttrFd(fd); err == nil && xattr&FsImmutableFlag != 0 {
		return true
	}

	return false
}

// IsAppendOnly checks if file has append only flag set.
func IsAppendOnly(fd *os.File) bool {
	if xattr, err := GetAttr(fd); err == nil && xattr&FsAppendFlag != 0 {
		return true
	}

	return false
}

// IsAppendOnlyFdPtr checks if file has append only flag set, using file descriptor as uintptr directly.
func IsAppendOnlyFdPtr(fd uintptr) bool {
	if xattr, err := GetAttrFdPtr(fd); err == nil && xattr&FsAppendFlag != 0 {
		return true
	}

	return false
}

// IsAppendOnlyFd checks if file has append only flag set, using file descriptor as int directly.
func IsAppendOnlyFd(fd int) bool {
	if xattr, err := GetAttrFd(fd); err == nil && xattr&FsAppendFlag != 0 {
		return true
	}

	return false
}

// IsNoDump checks if file has no dump flag set.
func IsNoDump(fd *os.File) bool {
	if xattr, err := GetAttr(fd); err == nil && xattr&FsNoDumpFlag != 0 {
		return true
	}

	return false
}

// IsNoDumpFdPtr checks if file has no dump flag set, using file descriptor as uintptr directly.
func IsNoDumpFdPtr(fd uintptr) bool {
	if xattr, err := GetAttrFdPtr(fd); err == nil && xattr&FsNoDumpFlag != 0 {
		return true
	}

	return false
}

// IsNoDumpFd checks if file has no dump flag set, using file descriptor as int directly.
func IsNoDumpFd(fd int) bool {
	if xattr, err := GetAttrFd(fd); err == nil && xattr&FsNoDumpFlag != 0 {
		return true
	}

	return false
}

// SetImmutable sets immutable flag on file.
func SetImmutable(fd *os.File) error {
	return SetAttr(fd, FsImmutableFlag)
}

// SetImmutableFdPtr sets immutable flag on file, using file descriptor as uintptr directly.
func SetImmutableFdPtr(fd uintptr) error {
	return SetAttrFdPtr(fd, FsImmutableFlag)
}

// SetImmutableFd sets immutable flag on file, using file descriptor as int directly.
func SetImmutableFd(fd int) error {
	return SetAttrFd(fd, FsImmutableFlag)
}

// SetAppendOnly sets append only flag on file.
func SetAppendOnly(fd *os.File) error {
	return SetAttr(fd, FsAppendFlag)
}

// SetAppendOnlyFdPtr sets append only flag on file, using file descriptor as uintptr directly.
func SetAppendOnlyFdPtr(fd uintptr) error {
	return SetAttrFdPtr(fd, FsAppendFlag)
}

// SetAppendOnlyFd sets append only flag on file, using file descriptor as int directly.
func SetAppendOnlyFd(fd int) error {
	return SetAttrFd(fd, FsAppendFlag)
}

// SetNoDump sets no dump flag on file.
func SetNoDump(fd *os.File) error {
	return SetAttr(fd, FsNoDumpFlag)
}

// SetNoDumpFdPtr sets no dump flag on file, using file descriptor as uintptr directly.
func SetNoDumpFdPtr(fd uintptr) error {
	return SetAttrFdPtr(fd, FsNoDumpFlag)
}

// SetNoDumpFd sets no dump flag on file, using file descriptor as int directly.
func SetNoDumpFd(fd int) error {
	return SetAttrFd(fd, FsNoDumpFlag)
}

// UnSetImmutable unsets immutable flag on file.
func UnSetImmutable(fd *os.File) error {
	xattr, err := GetAttr(fd)
	if err != nil {
		return err
	}

	return SetAttr(fd, xattr&^FsImmutableFlag)
}

// UnSetImmutableFdPtr unsets immutable flag on file, using file descriptor as uintptr directly.
func UnSetImmutableFdPtr(fd uintptr) error {
	xattr, err := GetAttrFdPtr(fd)
	if err != nil {
		return err
	}

	return SetAttrFdPtr(fd, xattr&^FsImmutableFlag)
}

// UnSetImmutableFd unsets immutable flag on file, using file descriptor as int directly.
func UnSetImmutableFd(fd int) error {
	xattr, err := GetAttrFd(fd)
	if err != nil {
		return err
	}

	return SetAttrFd(fd, xattr&^FsImmutableFlag)
}

// UnSetAppendOnly unsets append only flag on file.
func UnSetAppendOnly(fd *os.File) error {
	xattr, err := GetAttr(fd)
	if err != nil {
		return err
	}

	return SetAttr(fd, xattr&^FsAppendFlag)
}

// UnSetAppendOnlyFdPtr unsets append only flag on file, using file descriptor as uintptr directly.
func UnSetAppendOnlyFdPtr(fd uintptr) error {
	xattr, err := GetAttrFdPtr(fd)
	if err != nil {
		return err
	}

	return SetAttrFdPtr(fd, xattr&^FsAppendFlag)
}

// UnSetAppendOnlyFd unsets append only flag on file, using file descriptor as int directly.
func UnSetAppendOnlyFd(fd int) error {
	xattr, err := GetAttrFd(fd)
	if err != nil {
		return err
	}

	return SetAttrFd(fd, xattr&^FsAppendFlag)
}

// UnSetNoDump unsets no dump flag on file.
func UnSetNoDump(fd *os.File) error {
	xattr, err := GetAttr(fd)
	if err != nil {
		return err
	}

	return SetAttr(fd, xattr&^FsNoDumpFlag)
}

// UnSetNoDumpFdPtr unsets no dump flag on file, using file descriptor as uintptr directly.
func UnSetNoDumpFdPtr(fd uintptr) error {
	xattr, err := GetAttrFdPtr(fd)
	if err != nil {
		return err
	}

	return SetAttrFdPtr(fd, xattr&^FsNoDumpFlag)
}

// UnSetNoDumpFd unsets no dump flag on file, using file descriptor as int directly.
func UnSetNoDumpFd(fd int) error {
	xattr, err := GetAttrFd(fd)
	if err != nil {
		return err
	}

	return SetAttrFd(fd, xattr&^FsNoDumpFlag)
}
