//go:build linux
// +build linux

// Package linux provides low-level Linux system call wrappers for x64 architecture.
// These functions directly invoke kernel system calls using the syscall package.
// Use with caution as they bypass Go's standard library safety checks.
package linux

import (
	"syscall"
	"unsafe"
)

// SyscallNumber represents a Linux system call number for x64 architecture.
type SyscallNumber int

const (
	__x64_sys_read SyscallNumber = iota
	__x64_sys_write
	__x64_sys_open
	__x64_sys_close
	__x64_sys_newstat
	__x64_sys_newfstat
	__x64_sys_newlstat
	__x64_sys_poll
	__x64_sys_lseek
	__x64_sys_mmap
	__x64_sys_mprotect
	__x64_sys_munmap
	__x64_sys_brk
	__x64_sys_rt_sigaction
	__x64_sys_rt_sigprocmask
	__x64_sys_rt_sigreturn
	__x64_sys_ioctl
	__x64_sys_pread64
	__x64_sys_pwrite64
	__x64_sys_readv
	__x64_sys_writev
	__x64_sys_access
	__x64_sys_pipe
	__x64_sys_select
	__x64_sys_sched_yield
	__x64_sys_mremap
	__x64_sys_msync
	__x64_sys_mincore
	__x64_sys_madvise
	__x64_sys_shmget
	__x64_sys_shmat
	__x64_sys_shmctl
	__x64_sys_dup
	__x64_sys_dup2
	__x64_sys_pause
	__x64_sys_nanosleep
	__x64_sys_getitimer
	__x64_sys_alarm
	__x64_sys_setitimer
	__x64_sys_getpid
	__x64_sys_sendfile64
	__x64_sys_socket
	__x64_sys_connect
	__x64_sys_accept
	__x64_sys_sendto
	__x64_sys_recvfrom
	__x64_sys_sendmsg
	__x64_sys_recvmsg
	__x64_sys_shutdown
	__x64_sys_bind
	__x64_sys_listen
	__x64_sys_getsockname
	__x64_sys_getpeername
	__x64_sys_socketpair
	__x64_sys_setsockopt
	__x64_sys_getsockopt
	__x64_sys_clone
	__x64_sys_fork
	__x64_sys_vfork
	__x64_sys_execve
	__x64_sys_exit
	__x64_sys_wait4
	__x64_sys_kill
	__x64_sys_newuname
	__x64_sys_semget
	__x64_sys_semop
	__x64_sys_semctl
	__x64_sys_shmdt
	__x64_sys_msgget
	__x64_sys_msgsnd
	__x64_sys_msgrcv
	__x64_sys_msgctl
	__x64_sys_fcntl
	__x64_sys_flock
	__x64_sys_fsync
	__x64_sys_fdatasync
	__x64_sys_truncate
	__x64_sys_ftruncate
	__x64_sys_getdents
	__x64_sys_getcwd
	__x64_sys_chdir
	__x64_sys_fchdir
	__x64_sys_rename
	__x64_sys_mkdir
	__x64_sys_rmdir
	__x64_sys_creat
	__x64_sys_link
	__x64_sys_unlink
	__x64_sys_symlink
	__x64_sys_readlink
	__x64_sys_chmod
	__x64_sys_fchmod
	__x64_sys_chown
	__x64_sys_fchown
	__x64_sys_lchown
	__x64_sys_umask
	__x64_sys_gettimeofday
	__x64_sys_getrlimit
	__x64_sys_getrusage
	__x64_sys_sysinfo
	__x64_sys_times
	__x64_sys_ptrace
	__x64_sys_getuid
	__x64_sys_syslog
	__x64_sys_getgid
	__x64_sys_setuid
	__x64_sys_setgid
	__x64_sys_geteuid
	__x64_sys_getegid
	__x64_sys_setpgid
	__x64_sys_getppid
	__x64_sys_getpgrp
	__x64_sys_setsid
	__x64_sys_setreuid
	__x64_sys_setregid
	__x64_sys_getgroups
	__x64_sys_setgroups
	__x64_sys_setresuid
	__x64_sys_getresuid
	__x64_sys_setresgid
	__x64_sys_getresgid
	__x64_sys_getpgid
	__x64_sys_setfsuid
	__x64_sys_setfsgid
	__x64_sys_getsid
	__x64_sys_capget
	__x64_sys_capset
	__x64_sys_rt_sigpending
	__x64_sys_rt_sigtimedwait
	__x64_sys_rt_sigqueueinfo
	__x64_sys_rt_sigsuspend
	__x64_sys_sigaltstack
	__x64_sys_utime
	__x64_sys_mknod
	__x64_sys_personality
	__x64_sys_ustat
	__x64_sys_statfs
	__x64_sys_fstatfs
	__x64_sys_sysfs
	__x64_sys_getpriority
	__x64_sys_setpriority
	__x64_sys_sched_setparam
	__x64_sys_sched_getparam
	__x64_sys_sched_setscheduler
	__x64_sys_sched_getscheduler
	__x64_sys_sched_get_priority_max
	__x64_sys_sched_get_priority_min
	__x64_sys_sched_rr_get_interval
	__x64_sys_mlock
	__x64_sys_munlock
	__x64_sys_mlockall
	__x64_sys_munlockall
	__x64_sys_vhangup
	__x64_sys_modify_ldt
	__x64_sys_pivot_root
	__x64_sys_prctl
	__x64_sys_arch_prctl
	__x64_sys_adjtimex
	__x64_sys_setrlimit
	__x64_sys_chroot
	__x64_sys_sync
	__x64_sys_acct
	__x64_sys_settimeofday
	__x64_sys_mount
	__x64_sys_umount
	__x64_sys_swapon
	__x64_sys_swapoff
	__x64_sys_reboot
	__x64_sys_sethostname
	__x64_sys_setdomainname
	__x64_sys_iopl
	__x64_sys_ioperm
	__x64_sys_init_module
	__x64_sys_delete_module
	__x64_sys_quotactl
	__x64_sys_gettid
	__x64_sys_readahead
	__x64_sys_setxattr
	__x64_sys_lsetxattr
	__x64_sys_fsetxattr
	__x64_sys_getxattr
	__x64_sys_lgetxattr
	__x64_sys_fgetxattr
	__x64_sys_listxattr
	__x64_sys_llistxattr
	__x64_sys_flistxattr
	__x64_sys_removexattr
	__x64_sys_lremovexattr
	__x64_sys_fremovexattr
	__x64_sys_tkill
	__x64_sys_time
	__x64_sys_futex
	__x64_sys_sched_setaffinity
	__x64_sys_sched_getaffinity
	__x64_sys_io_setup
	__x64_sys_io_destroy
	__x64_sys_io_getevents
	__x64_sys_io_submit
	__x64_sys_io_cancel
	__x64_sys_epoll_create
	__x64_sys_remap_file_pages
	__x64_sys_getdents64
	__x64_sys_set_tid_address
	__x64_sys_restart_syscall
	__x64_sys_semtimedop
	__x64_sys_fadvise64
	__x64_sys_timer_create
	__x64_sys_timer_settime
	__x64_sys_timer_gettime
	__x64_sys_timer_getoverrun
	__x64_sys_timer_delete
	__x64_sys_clock_settime
	__x64_sys_clock_gettime
	__x64_sys_clock_getres
	__x64_sys_clock_nanosleep
	__x64_sys_exit_group
	__x64_sys_epoll_wait
	__x64_sys_epoll_ctl
	__x64_sys_tgkill
	__x64_sys_utimes
	__x64_sys_mbind
	__x64_sys_set_mempolicy
	__x64_sys_get_mempolicy
	__x64_sys_mq_open
	__x64_sys_mq_unlink
	__x64_sys_mq_timedsend
	__x64_sys_mq_timedreceive
	__x64_sys_mq_notify
	__x64_sys_mq_getsetattr
	__x64_sys_kexec_load
	__x64_sys_waitid
	__x64_sys_add_key
	__x64_sys_request_key
	__x64_sys_keyctl
	__x64_sys_ioprio_set
	__x64_sys_ioprio_get
	__x64_sys_inotify_init
	__x64_sys_inotify_add_watch
	__x64_sys_inotify_rm_watch
	__x64_sys_migrate_pages
	__x64_sys_openat
	__x64_sys_mkdirat
	__x64_sys_mknodat
	__x64_sys_fchownat
	__x64_sys_futimesat
	__x64_sys_newfstatat
	__x64_sys_unlinkat
	__x64_sys_renameat
	__x64_sys_linkat
	__x64_sys_symlinkat
	__x64_sys_readlinkat
	__x64_sys_fchmodat
	__x64_sys_faccessat
	__x64_sys_pselect6
	__x64_sys_ppoll
	__x64_sys_unshare
	__x64_sys_set_robust_list
	__x64_sys_get_robust_list
	__x64_sys_splice
	__x64_sys_tee
	__x64_sys_sync_file_range
	__x64_sys_vmsplice
	__x64_sys_move_pages
	__x64_sys_utimensat
	__x64_sys_epoll_pwait
	__x64_sys_signalfd
	__x64_sys_timerfd_create
	__x64_sys_eventfd
	__x64_sys_fallocate
	__x64_sys_timerfd_settime
	__x64_sys_timerfd_gettime
	__x64_sys_accept4
	__x64_sys_signalfd4
	__x64_sys_eventfd2
	__x64_sys_epoll_create1
	__x64_sys_dup3
	__x64_sys_pipe2
	__x64_sys_inotify_init1
	__x64_sys_preadv
	__x64_sys_pwritev
	__x64_sys_rt_tgsigqueueinfo
	__x64_sys_perf_event_open
	__x64_sys_recvmmsg
	__x64_sys_fanotify_init
	__x64_sys_fanotify_mark
	__x64_sys_prlimit64
	__x64_sys_name_to_handle_at
	__x64_sys_open_by_handle_at
	__x64_sys_clock_adjtime
	__x64_sys_syncfs
	__x64_sys_sendmmsg
	__x64_sys_setns
	__x64_sys_getcpu
	__x64_sys_process_vm_readv
	__x64_sys_process_vm_writev
	__x64_sys_kcmp
	__x64_sys_finit_module
	__x64_sys_sched_setattr
	__x64_sys_sched_getattr
	__x64_sys_renameat2
	__x64_sys_seccomp
	__x64_sys_getrandom
	__x64_sys_memfd_create
	__x64_sys_kexec_file_load
	__x64_sys_bpf
	__x64_sys_execveat
	__x64_sys_userfaultfd
	__x64_sys_membarrier
	__x64_sys_mlock2
	__x64_sys_copy_file_range
	__x64_sys_preadv2
	__x64_sys_pwritev2
	__x64_sys_pkey_mprotect
	__x64_sys_pkey_alloc
	__x64_sys_pkey_free
	__x64_sys_statx
	__x64_sys_io_pgetevents
	__x64_sys_rseq
	__x64_sys_uretprobe
	__x64_sys_pidfd_send_signal
	__x64_sys_io_uring_setup
	__x64_sys_io_uring_enter
	__x64_sys_io_uring_register
	__x64_sys_open_tree
	__x64_sys_move_mount
	__x64_sys_fsopen
	__x64_sys_fsconfig
	__x64_sys_fsmount
	__x64_sys_fspick
	__x64_sys_pidfd_open
	__x64_sys_clone3
	__x64_sys_close_range
	__x64_sys_openat2
	__x64_sys_pidfd_getfd
	__x64_sys_faccessat2
	__x64_sys_process_madvise
	__x64_sys_epoll_pwait2
	__x64_sys_mount_setattr
	__x64_sys_quotactl_fd
	__x64_sys_landlock_create_ruleset
	__x64_sys_landlock_add_rule
	__x64_sys_landlock_restrict_self
	__x64_sys_memfd_secret
	__x64_sys_process_mrelease
	__x64_sys_futex_waitv
	__x64_sys_set_mempolicy_home_node
	__x64_sys_cachestat
	__x64_sys_fchmodat2
	__x64_sys_map_shadow_stack
	__x64_sys_futex_wake
	__x64_sys_futex_wait
	__x64_sys_futex_requeue
	__x64_sys_statmount
	__x64_sys_listmount
	__x64_sys_lsm_get_self_attr
	__x64_sys_lsm_set_self_attr
	__x64_sys_lsm_list_modules
	__x64_sys_mseal
	__x64_sys_setxattrat
	__x64_sys_getxattrat
	__x64_sys_listxattrat
	__x64_sys_removexattrat
	__x64_sys_open_tree_attr
	__x64_sys_file_getattr
	__x64_sys_file_setattr
)

// Read reads up to len(buf) bytes from the file descriptor fd into buf.
// It returns the number of bytes read and any error encountered.
func Read(fd uint, buf []byte) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(
		uintptr(__x64_sys_read),
		uintptr(fd),
		uintptr(unsafe.Pointer(&buf[0])),
		uintptr(len(buf)),
	)
	return r1, err
}

// Write writes len(buf) bytes from buf to the file descriptor fd.
// It returns the number of bytes written and any error encountered.
func Write(fd uint, buf []byte) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(
		uintptr(__x64_sys_write),
		uintptr(fd),
		uintptr(unsafe.Pointer(&buf[0])),
		uintptr(len(buf)),
	)

	return r1, err
}

// Open opens the file specified by filename with the given flags and mode.
// It returns a file descriptor and any error encountered.
func Open(filename []byte, flags int, mode uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(
		uintptr(__x64_sys_open),
		uintptr(unsafe.Pointer(&filename[0])),
		uintptr(flags),
		uintptr(mode),
	)

	return r1, err
}

// Close closes the file descriptor fd.
// It returns 0 on success and any error encountered.
func Close(fd uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(
		uintptr(__x64_sys_close),
		uintptr(fd),
		0,
		0,
	)
	return r1, err
}

func Newstat(filename []byte, statbuf unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_newstat), uintptr(unsafe.Pointer(&filename[0])), uintptr(statbuf), 0)
	return r1, err
}

func Newfstat(fd uint, statbuf unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_newfstat), uintptr(fd), uintptr(statbuf), 0)
	return r1, err
}

func Newlstat(filename []byte, statbuf unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_newlstat), uintptr(unsafe.Pointer(&filename[0])), uintptr(statbuf), 0)
	return r1, err
}

func Poll(ufds unsafe.Pointer, nfds uint, timeout int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_poll), uintptr(ufds), uintptr(nfds), uintptr(timeout))
	return r1, err
}

func Lseek(fd uint, offset int64, whence uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_lseek), uintptr(fd), uintptr(offset), uintptr(whence))
	return r1, err
}

// Mmap maps length bytes starting at offset in the file specified by fd into memory.
// It returns the address of the mapped region and any error encountered.
func Mmap(addr uintptr, length uintptr, prot uintptr, flags uintptr, fd uintptr, offset uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_mmap), addr, length, prot, flags, fd, offset)
	return r1, err
}

func Mprotect(start uintptr, length uintptr, prot uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_mprotect), start, length, prot)
	return r1, err
}

func Munmap(addr uintptr, length uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_munmap), addr, length, 0)
	return r1, err
}

func Brk(brk uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_brk), brk, 0, 0)
	return r1, err
}

func RtSigaction(sig int, act unsafe.Pointer, oact unsafe.Pointer, sigsetsize uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_rt_sigaction), uintptr(sig), uintptr(act), uintptr(oact), sigsetsize, 0, 0)
	return r1, err
}

func RtSigprocmask(how int, nset unsafe.Pointer, oset unsafe.Pointer, sigsetsize uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_rt_sigprocmask), uintptr(how), uintptr(nset), uintptr(oset), sigsetsize, 0, 0)
	return r1, err
}

func RtSigreturn() (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_rt_sigreturn), 0, 0, 0)
	return r1, err
}

func Ioctl(fd uint, cmd uint, arg uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_ioctl), uintptr(fd), uintptr(cmd), arg)
	return r1, err
}

func Pread64(fd uint, buf unsafe.Pointer, count uintptr, pos int64) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_pread64), uintptr(fd), uintptr(buf), count, uintptr(pos), 0, 0)
	return r1, err
}

func Pwrite64(fd uint, buf unsafe.Pointer, count uintptr, pos int64) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_pwrite64), uintptr(fd), uintptr(buf), count, uintptr(pos), 0, 0)
	return r1, err
}

func Readv(fd uintptr, vec unsafe.Pointer, vlen uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_readv), fd, uintptr(vec), vlen)
	return r1, err
}

func Writev(fd uintptr, vec unsafe.Pointer, vlen uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_writev), fd, uintptr(vec), vlen)
	return r1, err
}

func Access(filename []byte, mode int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_access), uintptr(unsafe.Pointer(&filename[0])), uintptr(mode), 0)
	return r1, err
}

func Pipe(fildes unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_pipe), uintptr(fildes), 0, 0)
	return r1, err
}

func Select(n int, inp unsafe.Pointer, outp unsafe.Pointer, exp unsafe.Pointer, tvp unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_select), uintptr(n), uintptr(inp), uintptr(outp), uintptr(exp), uintptr(tvp), 0)
	return r1, err
}

func SchedYield() (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_sched_yield), 0, 0, 0)
	return r1, err
}

func Mremap(addr uintptr, oldLen uintptr, newLen uintptr, flags uintptr, newAddr uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_mremap), addr, oldLen, newLen, flags, newAddr, 0)
	return r1, err
}

func Msync(start uintptr, length uintptr, flags int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_msync), start, length, uintptr(flags))
	return r1, err
}

func Mincore(start uintptr, length uintptr, vec unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_mincore), start, length, uintptr(vec))
	return r1, err
}

func Madvise(start uintptr, lenIn uintptr, behavior int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_madvise), start, lenIn, uintptr(behavior))
	return r1, err
}

func Shmget(key int, size uintptr, shmflg int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_shmget), uintptr(key), size, uintptr(shmflg))
	return r1, err
}

func Shmat(shmid int, shmaddr unsafe.Pointer, shmflg int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_shmat), uintptr(shmid), uintptr(shmaddr), uintptr(shmflg))
	return r1, err
}

func Shmctl(shmid int, cmd int, buf unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_shmctl), uintptr(shmid), uintptr(cmd), uintptr(buf))
	return r1, err
}

func Dup(fildes uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_dup), uintptr(fildes), 0, 0)
	return r1, err
}

func Dup2(oldfd uint, newfd uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_dup2), uintptr(oldfd), uintptr(newfd), 0)
	return r1, err
}

func Pause() (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_pause), 0, 0, 0)
	return r1, err
}

func Nanosleep(rqtp unsafe.Pointer, rmtp unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_nanosleep), uintptr(rqtp), uintptr(rmtp), 0)
	return r1, err
}

func Getitimer(which int, value unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_getitimer), uintptr(which), uintptr(value), 0)
	return r1, err
}

func Alarm(seconds uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_alarm), uintptr(seconds), 0, 0)
	return r1, err
}

func Setitimer(which int, value unsafe.Pointer, ovalue unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_setitimer), uintptr(which), uintptr(value), uintptr(ovalue))
	return r1, err
}

// Getpid returns the process ID of the calling process.
func Getpid() (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_getpid), 0, 0, 0)
	return r1, err
}

func Sendfile64(outFd int, inFd int, offset unsafe.Pointer, count uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_sendfile64), uintptr(outFd), uintptr(inFd), uintptr(offset), count, 0, 0)
	return r1, err
}

func Socket(family int, typ int, protocol int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_socket), uintptr(family), uintptr(typ), uintptr(protocol))
	return r1, err
}

func Connect(fd int, uservaddr unsafe.Pointer, addrlen int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_connect), uintptr(fd), uintptr(uservaddr), uintptr(addrlen))
	return r1, err
}

func Accept(fd int, upeerSockaddr unsafe.Pointer, upeerAddrlen unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_accept), uintptr(fd), uintptr(upeerSockaddr), uintptr(upeerAddrlen))
	return r1, err
}

func Sendto(fd int, buff unsafe.Pointer, length uintptr, flags uint, addr unsafe.Pointer, addrLen int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_sendto), uintptr(fd), uintptr(buff), length, uintptr(flags), uintptr(addr), uintptr(addrLen))
	return r1, err
}

func Recvfrom(fd int, ubuf unsafe.Pointer, size uintptr, flags uint, addr unsafe.Pointer, addrLen unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_recvfrom), uintptr(fd), uintptr(ubuf), size, uintptr(flags), uintptr(addr), uintptr(addrLen))
	return r1, err
}

func Sendmsg(fd int, msg unsafe.Pointer, flags uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_sendmsg), uintptr(fd), uintptr(msg), uintptr(flags))
	return r1, err
}

func Recvmsg(fd int, msg unsafe.Pointer, flags uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_recvmsg), uintptr(fd), uintptr(msg), uintptr(flags))
	return r1, err
}

func Shutdown(fd int, how int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_shutdown), uintptr(fd), uintptr(how), 0)
	return r1, err
}

func Bind(fd int, umyaddr unsafe.Pointer, addrlen int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_bind), uintptr(fd), uintptr(umyaddr), uintptr(addrlen))
	return r1, err
}

func Listen(fd int, backlog int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_listen), uintptr(fd), uintptr(backlog), 0)
	return r1, err
}

func Getsockname(fd int, usockaddr unsafe.Pointer, usockaddrLen unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_getsockname), uintptr(fd), uintptr(usockaddr), uintptr(usockaddrLen))
	return r1, err
}

func Getpeername(fd int, usockaddr unsafe.Pointer, usockaddrLen unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_getpeername), uintptr(fd), uintptr(usockaddr), uintptr(usockaddrLen))
	return r1, err
}

func Socketpair(family int, typ int, protocol int, usockvec unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_socketpair), uintptr(family), uintptr(typ), uintptr(protocol), uintptr(usockvec), 0, 0)
	return r1, err
}

func Setsockopt(fd int, level int, optname int, optval unsafe.Pointer, optlen int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_setsockopt), uintptr(fd), uintptr(level), uintptr(optname), uintptr(optval), uintptr(optlen), 0)
	return r1, err
}

func Getsockopt(fd int, level int, optname int, optval unsafe.Pointer, optlen unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_getsockopt), uintptr(fd), uintptr(level), uintptr(optname), uintptr(optval), uintptr(optlen), 0)
	return r1, err
}

func Clone(cloneFlags uintptr, newsp uintptr, parentTidptr unsafe.Pointer, childTidptr unsafe.Pointer, tls uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_clone), cloneFlags, newsp, uintptr(parentTidptr), uintptr(childTidptr), tls, 0)
	return r1, err
}

// Fork creates a new process by duplicating the calling process.
// It returns the process ID of the child in the parent, and 0 in the child.
func Fork() (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_fork), 0, 0, 0)
	return r1, err
}

func Vfork() (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_vfork), 0, 0, 0)
	return r1, err
}

func Execve(filename []byte, argv unsafe.Pointer, envp unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_execve), uintptr(unsafe.Pointer(&filename[0])), uintptr(argv), uintptr(envp))
	return r1, err
}

func Exit(errorCode int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_exit), uintptr(errorCode), 0, 0)
	return r1, err
}

func Wait4(upid int, statAddr unsafe.Pointer, options int, ru unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_wait4), uintptr(upid), uintptr(statAddr), uintptr(options), uintptr(ru), 0, 0)
	return r1, err
}

func Kill(pid int, sig int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_kill), uintptr(pid), uintptr(sig), 0)
	return r1, err
}

func Newuname(name unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_newuname), uintptr(name), 0, 0)
	return r1, err
}

func Semget(key int, nsems int, semflg int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_semget), uintptr(key), uintptr(nsems), uintptr(semflg))
	return r1, err
}

func Semop(semid int, tsops unsafe.Pointer, nsops uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_semop), uintptr(semid), uintptr(tsops), uintptr(nsops))
	return r1, err
}

func Semctl(semid int, semnum int, cmd int, arg uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_semctl), uintptr(semid), uintptr(semnum), uintptr(cmd), arg, 0, 0)
	return r1, err
}

func Shmdt(shmaddr unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_shmdt), uintptr(shmaddr), 0, 0)
	return r1, err
}

func Msgget(key int, msgflg int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_msgget), uintptr(key), uintptr(msgflg), 0)
	return r1, err
}

func Msgsnd(msqid int, msgp unsafe.Pointer, msgsz uintptr, msgflg int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_msgsnd), uintptr(msqid), uintptr(msgp), msgsz, uintptr(msgflg), 0, 0)
	return r1, err
}

func Msgrcv(msqid int, msgp unsafe.Pointer, msgsz uintptr, msgtyp int64, msgflg int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_msgrcv), uintptr(msqid), uintptr(msgp), msgsz, uintptr(msgtyp), uintptr(msgflg), 0)
	return r1, err
}

func Msgctl(msqid int, cmd int, buf unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_msgctl), uintptr(msqid), uintptr(cmd), uintptr(buf))
	return r1, err
}

func Fcntl(fd uint, cmd uint, arg uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_fcntl), uintptr(fd), uintptr(cmd), arg)
	return r1, err
}

func Flock(fd uint, cmd uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_flock), uintptr(fd), uintptr(cmd), 0)
	return r1, err
}

func Fsync(fd uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_fsync), uintptr(fd), 0, 0)
	return r1, err
}

func Fdatasync(fd uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_fdatasync), uintptr(fd), 0, 0)
	return r1, err
}

func Truncate(path []byte, length int64) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_truncate), uintptr(unsafe.Pointer(&path[0])), uintptr(length), 0)
	return r1, err
}

func Ftruncate(fd uint, length int64) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_ftruncate), uintptr(fd), uintptr(length), 0)
	return r1, err
}

func Getdents(fd uint, dirent unsafe.Pointer, count uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_getdents), uintptr(fd), uintptr(dirent), uintptr(count))
	return r1, err
}

func Getcwd(buf unsafe.Pointer, size uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_getcwd), uintptr(buf), size, 0)
	return r1, err
}

func Chdir(filename []byte) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_chdir), uintptr(unsafe.Pointer(&filename[0])), 0, 0)
	return r1, err
}

func Fchdir(fd uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_fchdir), uintptr(fd), 0, 0)
	return r1, err
}

func Rename(oldname []byte, newname []byte) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_rename), uintptr(unsafe.Pointer(&oldname[0])), uintptr(unsafe.Pointer(&newname[0])), 0)
	return r1, err
}

func Mkdir(pathname []byte, mode uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_mkdir), uintptr(unsafe.Pointer(&pathname[0])), uintptr(mode), 0)
	return r1, err
}

func Rmdir(pathname []byte) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_rmdir), uintptr(unsafe.Pointer(&pathname[0])), 0, 0)
	return r1, err
}

func Creat(pathname []byte, mode uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_creat), uintptr(unsafe.Pointer(&pathname[0])), uintptr(mode), 0)
	return r1, err
}

func Link(oldname []byte, newname []byte) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_link), uintptr(unsafe.Pointer(&oldname[0])), uintptr(unsafe.Pointer(&newname[0])), 0)
	return r1, err
}

func Unlink(pathname []byte) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_unlink), uintptr(unsafe.Pointer(&pathname[0])), 0, 0)
	return r1, err
}

func Symlink(oldname []byte, newname []byte) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_symlink), uintptr(unsafe.Pointer(&oldname[0])), uintptr(unsafe.Pointer(&newname[0])), 0)
	return r1, err
}

func Readlink(path []byte, buf unsafe.Pointer, bufsiz int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_readlink), uintptr(unsafe.Pointer(&path[0])), uintptr(buf), uintptr(bufsiz))
	return r1, err
}

func Chmod(filename []byte, mode uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_chmod), uintptr(unsafe.Pointer(&filename[0])), uintptr(mode), 0)
	return r1, err
}

func Fchmod(fd uint, mode uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_fchmod), uintptr(fd), uintptr(mode), 0)
	return r1, err
}

func Chown(filename []byte, user uint, group uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_chown), uintptr(unsafe.Pointer(&filename[0])), uintptr(user), uintptr(group))
	return r1, err
}

func Fchown(fd uint, user uint, group uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_fchown), uintptr(fd), uintptr(user), uintptr(group))
	return r1, err
}

func Lchown(filename []byte, user uint, group uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_lchown), uintptr(unsafe.Pointer(&filename[0])), uintptr(user), uintptr(group))
	return r1, err
}

func Umask(mask int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_umask), uintptr(mask), 0, 0)
	return r1, err
}

func Gettimeofday(tv unsafe.Pointer, tz unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_gettimeofday), uintptr(tv), uintptr(tz), 0)
	return r1, err
}

func Getrlimit(resource uint, rlim unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_getrlimit), uintptr(resource), uintptr(rlim), 0)
	return r1, err
}

func Getrusage(who int, ru unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_getrusage), uintptr(who), uintptr(ru), 0)
	return r1, err
}

func Sysinfo(info unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_sysinfo), uintptr(info), 0, 0)
	return r1, err
}

func Times(tbuf unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_times), uintptr(tbuf), 0, 0)
	return r1, err
}

func Ptrace(request int64, pid int64, addr uintptr, data uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_ptrace), uintptr(request), uintptr(pid), addr, data, 0, 0)
	return r1, err
}

func Getuid() (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_getuid), 0, 0, 0)
	return r1, err
}

func Syslog(typ int, buf unsafe.Pointer, length int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_syslog), uintptr(typ), uintptr(buf), uintptr(length))
	return r1, err
}

func Getgid() (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_getgid), 0, 0, 0)
	return r1, err
}

func Setuid(uid uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_setuid), uintptr(uid), 0, 0)
	return r1, err
}

func Setgid(gid uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_setgid), uintptr(gid), 0, 0)
	return r1, err
}

func Geteuid() (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_geteuid), 0, 0, 0)
	return r1, err
}

func Getegid() (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_getegid), 0, 0, 0)
	return r1, err
}

func Setpgid(pid int, pgid int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_setpgid), uintptr(pid), uintptr(pgid), 0)
	return r1, err
}

func Getppid() (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_getppid), 0, 0, 0)
	return r1, err
}

func Getpgrp() (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_getpgrp), 0, 0, 0)
	return r1, err
}

func Setsid() (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_setsid), 0, 0, 0)
	return r1, err
}

func Setreuid(ruid uint, euid uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_setreuid), uintptr(ruid), uintptr(euid), 0)
	return r1, err
}

func Setregid(rgid uint, egid uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_setregid), uintptr(rgid), uintptr(egid), 0)
	return r1, err
}

func Getgroups(gidsetsize int, grouplist unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_getgroups), uintptr(gidsetsize), uintptr(grouplist), 0)
	return r1, err
}

func Setgroups(gidsetsize int, grouplist unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_setgroups), uintptr(gidsetsize), uintptr(grouplist), 0)
	return r1, err
}

func Setresuid(ruid uint, euid uint, suid uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_setresuid), uintptr(ruid), uintptr(euid), uintptr(suid))
	return r1, err
}

func Getresuid(ruidp unsafe.Pointer, euidp unsafe.Pointer, suidp unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_getresuid), uintptr(ruidp), uintptr(euidp), uintptr(suidp))
	return r1, err
}

func Setresgid(rgid uint, egid uint, sgid uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_setresgid), uintptr(rgid), uintptr(egid), uintptr(sgid))
	return r1, err
}

func Getresgid(rgidp unsafe.Pointer, egidp unsafe.Pointer, sgidp unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_getresgid), uintptr(rgidp), uintptr(egidp), uintptr(sgidp))
	return r1, err
}

func Getpgid(pid int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_getpgid), uintptr(pid), 0, 0)
	return r1, err
}

func Setfsuid(uid uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_setfsuid), uintptr(uid), 0, 0)
	return r1, err
}

func Setfsgid(gid uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_setfsgid), uintptr(gid), 0, 0)
	return r1, err
}

func Getsid(pid int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_getsid), uintptr(pid), 0, 0)
	return r1, err
}

func Capget(header unsafe.Pointer, dataptr unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_capget), uintptr(header), uintptr(dataptr), 0)
	return r1, err
}

func Capset(header unsafe.Pointer, data unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_capset), uintptr(header), uintptr(data), 0)
	return r1, err
}

func RtSigpending(uset unsafe.Pointer, sigsetsize uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_rt_sigpending), uintptr(uset), sigsetsize, 0)
	return r1, err
}

func RtSigtimedwait(uthese unsafe.Pointer, uinfo unsafe.Pointer, uts unsafe.Pointer, sigsetsize uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_rt_sigtimedwait), uintptr(uthese), uintptr(uinfo), uintptr(uts), sigsetsize, 0, 0)
	return r1, err
}

func RtSigqueueinfo(pid int, sig int, uinfo unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_rt_sigqueueinfo), uintptr(pid), uintptr(sig), uintptr(uinfo))
	return r1, err
}

func RtSigsuspend(unewset unsafe.Pointer, sigsetsize uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_rt_sigsuspend), uintptr(unewset), sigsetsize, 0)
	return r1, err
}

func Sigaltstack(uss unsafe.Pointer, uoss unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_sigaltstack), uintptr(uss), uintptr(uoss), 0)
	return r1, err
}

func Utime(filename []byte, times unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_utime), uintptr(unsafe.Pointer(&filename[0])), uintptr(times), 0)
	return r1, err
}

func Mknod(filename []byte, mode uint, dev uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_mknod), uintptr(unsafe.Pointer(&filename[0])), uintptr(mode), uintptr(dev))
	return r1, err
}

func Personality(personality uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_personality), uintptr(personality), 0, 0)
	return r1, err
}

func Ustat(dev uint, ubuf unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_ustat), uintptr(dev), uintptr(ubuf), 0)
	return r1, err
}

func Statfs(pathname []byte, buf unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_statfs), uintptr(unsafe.Pointer(&pathname[0])), uintptr(buf), 0)
	return r1, err
}

func Fstatfs(fd uint, buf unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_fstatfs), uintptr(fd), uintptr(buf), 0)
	return r1, err
}

func Sysfs(option int, arg1 uintptr, arg2 uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_sysfs), uintptr(option), arg1, arg2)
	return r1, err
}

func Getpriority(which int, who int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_getpriority), uintptr(which), uintptr(who), 0)
	return r1, err
}

func Setpriority(which int, who int, niceval int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_setpriority), uintptr(which), uintptr(who), uintptr(niceval))
	return r1, err
}

func SchedSetparam(pid int, param unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_sched_setparam), uintptr(pid), uintptr(param), 0)
	return r1, err
}

func SchedGetparam(pid int, param unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_sched_getparam), uintptr(pid), uintptr(param), 0)
	return r1, err
}

func SchedSetscheduler(pid int, policy int, param unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_sched_setscheduler), uintptr(pid), uintptr(policy), uintptr(param))
	return r1, err
}

func SchedGetscheduler(pid int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_sched_getscheduler), uintptr(pid), 0, 0)
	return r1, err
}

func SchedGetPriorityMax(policy int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_sched_get_priority_max), uintptr(policy), 0, 0)
	return r1, err
}

func SchedGetPriorityMin(policy int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_sched_get_priority_min), uintptr(policy), 0, 0)
	return r1, err
}

func SchedRrGetInterval(pid int, interval unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_sched_rr_get_interval), uintptr(pid), uintptr(interval), 0)
	return r1, err
}

func Mlock(start uintptr, length uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_mlock), start, length, 0)
	return r1, err
}

func Munlock(start uintptr, length uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_munlock), start, length, 0)
	return r1, err
}

func Mlockall(flags int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_mlockall), uintptr(flags), 0, 0)
	return r1, err
}

func Munlockall() (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_munlockall), 0, 0, 0)
	return r1, err
}

func Vhangup() (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_vhangup), 0, 0, 0)
	return r1, err
}

func ModifyLdt(func_ int, ptr unsafe.Pointer, bytecount uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_modify_ldt), uintptr(func_), uintptr(ptr), uintptr(bytecount))
	return r1, err
}

func PivotRoot(newRoot []byte, putOld []byte) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_pivot_root), uintptr(unsafe.Pointer(&newRoot[0])), uintptr(unsafe.Pointer(&putOld[0])), 0)
	return r1, err
}

func Prctl(option int, arg2 uintptr, arg3 uintptr, arg4 uintptr, arg5 uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_prctl), uintptr(option), arg2, arg3, arg4, arg5, 0)
	return r1, err
}

func ArchPrctl(option int, arg2 uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_arch_prctl), uintptr(option), arg2, 0)
	return r1, err
}

func Adjtimex(txc_p unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_adjtimex), uintptr(txc_p), 0, 0)
	return r1, err
}

func Setrlimit(resource uint, rlim unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_setrlimit), uintptr(resource), uintptr(rlim), 0)
	return r1, err
}

func Chroot(filename []byte) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_chroot), uintptr(unsafe.Pointer(&filename[0])), 0, 0)
	return r1, err
}

func Sync() (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_sync), 0, 0, 0)
	return r1, err
}

func Acct(name []byte) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_acct), uintptr(unsafe.Pointer(&name[0])), 0, 0)
	return r1, err
}

func Settimeofday(tv unsafe.Pointer, tz unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_settimeofday), uintptr(tv), uintptr(tz), 0)
	return r1, err
}

func Mount(devName []byte, dirName []byte, typ []byte, flags uintptr, data unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_mount), uintptr(unsafe.Pointer(&devName[0])), uintptr(unsafe.Pointer(&dirName[0])), uintptr(unsafe.Pointer(&typ[0])), flags, uintptr(data), 0)
	return r1, err
}

func Umount(name []byte, flags int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_umount), uintptr(unsafe.Pointer(&name[0])), uintptr(flags), 0)
	return r1, err
}

func Swapon(specialfile []byte, swapFlags int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_swapon), uintptr(unsafe.Pointer(&specialfile[0])), uintptr(swapFlags), 0)
	return r1, err
}

func Swapoff(specialfile []byte) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_swapoff), uintptr(unsafe.Pointer(&specialfile[0])), 0, 0)
	return r1, err
}

func Reboot(magic1 int, magic2 int, cmd uint, arg unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_reboot), uintptr(magic1), uintptr(magic2), uintptr(cmd), uintptr(arg), 0, 0)
	return r1, err
}

func Sethostname(name []byte, length int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_sethostname), uintptr(unsafe.Pointer(&name[0])), uintptr(length), 0)
	return r1, err
}

func Setdomainname(name []byte, length int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_setdomainname), uintptr(unsafe.Pointer(&name[0])), uintptr(length), 0)
	return r1, err
}

func Iopl(level uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_iopl), uintptr(level), 0, 0)
	return r1, err
}

func Ioperm(from uintptr, num uint, turnOn int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_ioperm), from, uintptr(num), uintptr(turnOn))
	return r1, err
}

func InitModule(umod unsafe.Pointer, length uintptr, uargs []byte) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_init_module), uintptr(umod), length, uintptr(unsafe.Pointer(&uargs[0])))
	return r1, err
}

func DeleteModule(nameUser []byte, flags uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_delete_module), uintptr(unsafe.Pointer(&nameUser[0])), uintptr(flags), 0)
	return r1, err
}

func Quotactl(cmd uint, special []byte, id int, addr unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_quotactl), uintptr(cmd), uintptr(unsafe.Pointer(&special[0])), uintptr(id), uintptr(addr), 0, 0)
	return r1, err
}

func Gettid() (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_gettid), 0, 0, 0)
	return r1, err
}

func Readahead(fd int, offset int64, count uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_readahead), uintptr(fd), uintptr(offset), count)
	return r1, err
}

func Setxattr(pathname []byte, name []byte, value unsafe.Pointer, size uintptr, flags int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_setxattr), uintptr(unsafe.Pointer(&pathname[0])), uintptr(unsafe.Pointer(&name[0])), uintptr(value), size, uintptr(flags), 0)
	return r1, err
}

func Lsetxattr(pathname []byte, name []byte, value unsafe.Pointer, size uintptr, flags int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_lsetxattr), uintptr(unsafe.Pointer(&pathname[0])), uintptr(unsafe.Pointer(&name[0])), uintptr(value), size, uintptr(flags), 0)
	return r1, err
}

func Fsetxattr(fd int, name []byte, value unsafe.Pointer, size uintptr, flags int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_fsetxattr), uintptr(fd), uintptr(unsafe.Pointer(&name[0])), uintptr(value), size, uintptr(flags), 0)
	return r1, err
}

func Getxattr(pathname []byte, name []byte, value unsafe.Pointer, size uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_getxattr), uintptr(unsafe.Pointer(&pathname[0])), uintptr(unsafe.Pointer(&name[0])), uintptr(value), size, 0, 0)
	return r1, err
}

func Lgetxattr(pathname []byte, name []byte, value unsafe.Pointer, size uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_lgetxattr), uintptr(unsafe.Pointer(&pathname[0])), uintptr(unsafe.Pointer(&name[0])), uintptr(value), size, 0, 0)
	return r1, err
}

func Fgetxattr(fd int, name []byte, value unsafe.Pointer, size uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_fgetxattr), uintptr(fd), uintptr(unsafe.Pointer(&name[0])), uintptr(value), size, 0, 0)
	return r1, err
}

func Listxattr(pathname []byte, list unsafe.Pointer, size uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_listxattr), uintptr(unsafe.Pointer(&pathname[0])), uintptr(list), size)
	return r1, err
}

func Llistxattr(pathname []byte, list unsafe.Pointer, size uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_llistxattr), uintptr(unsafe.Pointer(&pathname[0])), uintptr(list), size)
	return r1, err
}

func Flistxattr(fd int, list unsafe.Pointer, size uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_flistxattr), uintptr(fd), uintptr(list), size)
	return r1, err
}

func Removexattr(pathname []byte, name []byte) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_removexattr), uintptr(unsafe.Pointer(&pathname[0])), uintptr(unsafe.Pointer(&name[0])), 0)
	return r1, err
}

func Lremovexattr(pathname []byte, name []byte) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_lremovexattr), uintptr(unsafe.Pointer(&pathname[0])), uintptr(unsafe.Pointer(&name[0])), 0)
	return r1, err
}

func Fremovexattr(fd int, name []byte) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_fremovexattr), uintptr(fd), uintptr(unsafe.Pointer(&name[0])), 0)
	return r1, err
}

func Tkill(pid int, sig int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_tkill), uintptr(pid), uintptr(sig), 0)
	return r1, err
}

func Time(tloc unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_time), uintptr(tloc), 0, 0)
	return r1, err
}

func Futex(uaddr unsafe.Pointer, op int, val int, utime unsafe.Pointer, uaddr2 unsafe.Pointer, val3 int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_futex), uintptr(uaddr), uintptr(op), uintptr(val), uintptr(utime), uintptr(uaddr2), uintptr(val3))
	return r1, err
}

func SchedSetaffinity(pid int, length uint, userMaskPtr unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_sched_setaffinity), uintptr(pid), uintptr(length), uintptr(userMaskPtr))
	return r1, err
}

func SchedGetaffinity(pid int, length uint, userMaskPtr unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_sched_getaffinity), uintptr(pid), uintptr(length), uintptr(userMaskPtr))
	return r1, err
}

func IoSetup(nrEvents uintptr, ctxp unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_io_setup), nrEvents, uintptr(ctxp), 0)
	return r1, err
}

func IoDestroy(ctx uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_io_destroy), ctx, 0, 0)
	return r1, err
}

func IoGetevents(ctxId uintptr, minNr int64, nr int64, events unsafe.Pointer, timeout unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_io_getevents), ctxId, uintptr(minNr), uintptr(nr), uintptr(events), uintptr(timeout), 0)
	return r1, err
}

func IoSubmit(ctxId uintptr, nr int64, iocbpp unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_io_submit), ctxId, uintptr(nr), uintptr(iocbpp))
	return r1, err
}

func IoCancel(ctxId uintptr, iocb unsafe.Pointer, result unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_io_cancel), ctxId, uintptr(iocb), uintptr(result))
	return r1, err
}

func EpollCreate(size int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_epoll_create), uintptr(size), 0, 0)
	return r1, err
}

func RemapFilePages(start uintptr, size uintptr, prot uintptr, pgoff uintptr, flags uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_remap_file_pages), start, size, prot, pgoff, flags, 0)
	return r1, err
}

func Getdents64(fd uint, dirent unsafe.Pointer, count uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_getdents64), uintptr(fd), uintptr(dirent), uintptr(count))
	return r1, err
}

func SetTidAddress(tidptr unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_set_tid_address), uintptr(tidptr), 0, 0)
	return r1, err
}

func RestartSyscall() (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_restart_syscall), 0, 0, 0)
	return r1, err
}

func Semtimedop(semid int, tsops unsafe.Pointer, nsops uint, timeout unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_semtimedop), uintptr(semid), uintptr(tsops), uintptr(nsops), uintptr(timeout), 0, 0)
	return r1, err
}

func Fadvise64(fd int, offset int64, length int64, advice int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_fadvise64), uintptr(fd), uintptr(offset), uintptr(length), uintptr(advice), 0, 0)
	return r1, err
}

func TimerCreate(whichClock int, timerEventSpec unsafe.Pointer, createdTimerId unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_timer_create), uintptr(whichClock), uintptr(timerEventSpec), uintptr(createdTimerId))
	return r1, err
}

func TimerSettime(timerId uintptr, flags int, newSetting unsafe.Pointer, oldSetting unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_timer_settime), timerId, uintptr(flags), uintptr(newSetting), uintptr(oldSetting), 0, 0)
	return r1, err
}

func TimerGettime(timerId uintptr, setting unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_timer_gettime), timerId, uintptr(setting), 0)
	return r1, err
}

func TimerGetoverrun(timerId uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_timer_getoverrun), timerId, 0, 0)
	return r1, err
}

func TimerDelete(timerId uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_timer_delete), timerId, 0, 0)
	return r1, err
}

func ClockSettime(whichClock int, tp unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_clock_settime), uintptr(whichClock), uintptr(tp), 0)
	return r1, err
}

func ClockGettime(whichClock int, tp unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_clock_gettime), uintptr(whichClock), uintptr(tp), 0)
	return r1, err
}

func ClockGetres(whichClock int, tp unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_clock_getres), uintptr(whichClock), uintptr(tp), 0)
	return r1, err
}

func ClockNanosleep(whichClock int, flags int, rqtp unsafe.Pointer, rmtp unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_clock_nanosleep), uintptr(whichClock), uintptr(flags), uintptr(rqtp), uintptr(rmtp), 0, 0)
	return r1, err
}

func ExitGroup(errorCode int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_exit_group), uintptr(errorCode), 0, 0)
	return r1, err
}

func EpollWait(epfd int, events unsafe.Pointer, maxevents int, timeout int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_epoll_wait), uintptr(epfd), uintptr(events), uintptr(maxevents), uintptr(timeout), 0, 0)
	return r1, err
}

func EpollCtl(epfd int, op int, fd int, event unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_epoll_ctl), uintptr(epfd), uintptr(op), uintptr(fd), uintptr(event), 0, 0)
	return r1, err
}

func Tgkill(tgid int, pid int, sig int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_tgkill), uintptr(tgid), uintptr(pid), uintptr(sig))
	return r1, err
}

func Utimes(filename []byte, utimes unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_utimes), uintptr(unsafe.Pointer(&filename[0])), uintptr(utimes), 0)
	return r1, err
}

func Mbind(start uintptr, length uintptr, mode uintptr, nmask unsafe.Pointer, maxnode uintptr, flags uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_mbind), start, length, mode, uintptr(nmask), maxnode, flags)
	return r1, err
}

func SetMempolicy(mode int, nmask unsafe.Pointer, maxnode uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_set_mempolicy), uintptr(mode), uintptr(nmask), maxnode)
	return r1, err
}

func GetMempolicy(policy unsafe.Pointer, nmask unsafe.Pointer, maxnode uintptr, addr uintptr, flags uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_get_mempolicy), uintptr(policy), uintptr(nmask), maxnode, addr, flags, 0)
	return r1, err
}

func MqOpen(uName []byte, oflag int, mode uint, uAttr unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_mq_open), uintptr(unsafe.Pointer(&uName[0])), uintptr(oflag), uintptr(mode), uintptr(uAttr), 0, 0)
	return r1, err
}

func MqUnlink(uName []byte) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_mq_unlink), uintptr(unsafe.Pointer(&uName[0])), 0, 0)
	return r1, err
}

func MqTimedsend(mqdes uintptr, uMsgPtr unsafe.Pointer, msgLen uintptr, msgPrio uint, uAbsTimeout unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_mq_timedsend), mqdes, uintptr(uMsgPtr), msgLen, uintptr(msgPrio), uintptr(uAbsTimeout), 0)
	return r1, err
}

func MqTimedreceive(mqdes uintptr, uMsgPtr unsafe.Pointer, msgLen uintptr, uMsgPrio unsafe.Pointer, uAbsTimeout unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_mq_timedreceive), mqdes, uintptr(uMsgPtr), msgLen, uintptr(uMsgPrio), uintptr(uAbsTimeout), 0)
	return r1, err
}

func MqNotify(mqdes uintptr, uNotification unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_mq_notify), mqdes, uintptr(uNotification), 0)
	return r1, err
}

func MqGetsetattr(mqdes uintptr, uMqstat unsafe.Pointer, uOmqstat unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_mq_getsetattr), mqdes, uintptr(uMqstat), uintptr(uOmqstat))
	return r1, err
}

func KexecLoad(entry uintptr, nrSegments uintptr, segments unsafe.Pointer, flags uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_kexec_load), entry, nrSegments, uintptr(segments), flags, 0, 0)
	return r1, err
}

func Waitid(which int, upid int, infop unsafe.Pointer, options int, ru unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_waitid), uintptr(which), uintptr(upid), uintptr(infop), uintptr(options), uintptr(ru), 0)
	return r1, err
}

func AddKey(_type []byte, _description []byte, _payload unsafe.Pointer, plen uintptr, ringid int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_add_key), uintptr(unsafe.Pointer(&_type[0])), uintptr(unsafe.Pointer(&_description[0])), uintptr(_payload), plen, uintptr(ringid), 0)
	return r1, err
}

func RequestKey(_type []byte, _description []byte, _calloutInfo []byte, destringid int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_request_key), uintptr(unsafe.Pointer(&_type[0])), uintptr(unsafe.Pointer(&_description[0])), uintptr(unsafe.Pointer(&_calloutInfo[0])), uintptr(destringid), 0, 0)
	return r1, err
}

func Keyctl(option int, arg2 uintptr, arg3 uintptr, arg4 uintptr, arg5 uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_keyctl), uintptr(option), arg2, arg3, arg4, arg5, 0)
	return r1, err
}

func IoprioSet(which int, who int, ioprio int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_ioprio_set), uintptr(which), uintptr(who), uintptr(ioprio))
	return r1, err
}

func IoprioGet(which int, who int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_ioprio_get), uintptr(which), uintptr(who), 0)
	return r1, err
}

func InotifyInit() (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_inotify_init), 0, 0, 0)
	return r1, err
}

func InotifyAddWatch(fd int, pathname []byte, mask uint32) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_inotify_add_watch), uintptr(fd), uintptr(unsafe.Pointer(&pathname[0])), uintptr(mask))
	return r1, err
}

func InotifyRmWatch(fd int, wd int32) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_inotify_rm_watch), uintptr(fd), uintptr(wd), 0)
	return r1, err
}

func MigratePages(pid int, maxnode uintptr, oldNodes unsafe.Pointer, newNodes unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_migrate_pages), uintptr(pid), maxnode, uintptr(oldNodes), uintptr(newNodes), 0, 0)
	return r1, err
}

func Openat(dfd int, filename []byte, flags int, mode uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_openat), uintptr(dfd), uintptr(unsafe.Pointer(&filename[0])), uintptr(flags), uintptr(mode), 0, 0)
	return r1, err
}

func Mkdirat(dfd int, pathname []byte, mode uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_mkdirat), uintptr(dfd), uintptr(unsafe.Pointer(&pathname[0])), uintptr(mode))
	return r1, err
}

func Mknodat(dfd int, filename []byte, mode uint, dev uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_mknodat), uintptr(dfd), uintptr(unsafe.Pointer(&filename[0])), uintptr(mode), uintptr(dev), 0, 0)
	return r1, err
}

func Fchownat(dfd int, filename []byte, user uint, group uint, flag int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_fchownat), uintptr(dfd), uintptr(unsafe.Pointer(&filename[0])), uintptr(user), uintptr(group), uintptr(flag), 0)
	return r1, err
}

func Futimesat(dfd int, filename []byte, utimes unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_futimesat), uintptr(dfd), uintptr(unsafe.Pointer(&filename[0])), uintptr(utimes))
	return r1, err
}

func Newfstatat(dfd int, filename []byte, statbuf unsafe.Pointer, flag int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_newfstatat), uintptr(dfd), uintptr(unsafe.Pointer(&filename[0])), uintptr(statbuf), uintptr(flag), 0, 0)
	return r1, err
}

func Unlinkat(dfd int, pathname []byte, flag int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_unlinkat), uintptr(dfd), uintptr(unsafe.Pointer(&pathname[0])), uintptr(flag))
	return r1, err
}

func Renameat(olddfd int, oldname []byte, newdfd int, newname []byte) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_renameat), uintptr(olddfd), uintptr(unsafe.Pointer(&oldname[0])), uintptr(newdfd), uintptr(unsafe.Pointer(&newname[0])), 0, 0)
	return r1, err
}

func Linkat(olddfd int, oldname []byte, newdfd int, newname []byte, flags int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_linkat), uintptr(olddfd), uintptr(unsafe.Pointer(&oldname[0])), uintptr(newdfd), uintptr(unsafe.Pointer(&newname[0])), uintptr(flags), 0)
	return r1, err
}

func Symlinkat(oldname []byte, newdfd int, newname []byte) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_symlinkat), uintptr(unsafe.Pointer(&oldname[0])), uintptr(newdfd), uintptr(unsafe.Pointer(&newname[0])))
	return r1, err
}

func Readlinkat(dfd int, pathname []byte, buf unsafe.Pointer, bufsiz int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_readlinkat), uintptr(dfd), uintptr(unsafe.Pointer(&pathname[0])), uintptr(buf), uintptr(bufsiz), 0, 0)
	return r1, err
}

func Fchmodat(dfd int, filename []byte, mode uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_fchmodat), uintptr(dfd), uintptr(unsafe.Pointer(&filename[0])), uintptr(mode))
	return r1, err
}

func Faccessat(dfd int, filename []byte, mode int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_faccessat), uintptr(dfd), uintptr(unsafe.Pointer(&filename[0])), uintptr(mode))
	return r1, err
}

func Pselect6(n int, inp unsafe.Pointer, outp unsafe.Pointer, exp unsafe.Pointer, tsp unsafe.Pointer, sig unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_pselect6), uintptr(n), uintptr(inp), uintptr(outp), uintptr(exp), uintptr(tsp), uintptr(sig))
	return r1, err
}

func Ppoll(ufds unsafe.Pointer, nfds uintptr, tsp unsafe.Pointer, sigmask unsafe.Pointer, sigsetsize uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_ppoll), uintptr(ufds), nfds, uintptr(tsp), uintptr(sigmask), sigsetsize, 0)
	return r1, err
}

func Unshare(unshareFlags uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_unshare), unshareFlags, 0, 0)
	return r1, err
}

func SetRobustList(head unsafe.Pointer, length uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_set_robust_list), uintptr(head), length, 0)
	return r1, err
}

func GetRobustList(pid int, headPtr unsafe.Pointer, lenPtr unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_get_robust_list), uintptr(pid), uintptr(headPtr), uintptr(lenPtr))
	return r1, err
}

func Splice(fdIn int, offIn unsafe.Pointer, fdOut int, offOut unsafe.Pointer, length uintptr, flags uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_splice), uintptr(fdIn), uintptr(offIn), uintptr(fdOut), uintptr(offOut), length, uintptr(flags))
	return r1, err
}

func Tee(fdin int, fdout int, length uintptr, flags uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_tee), uintptr(fdin), uintptr(fdout), length, uintptr(flags), 0, 0)
	return r1, err
}

func SyncFileRange(fd int, offset int64, nbytes int64, flags uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_sync_file_range), uintptr(fd), uintptr(offset), uintptr(nbytes), uintptr(flags), 0, 0)
	return r1, err
}

func Vmsplice(fd int, uiov unsafe.Pointer, nrSegs uintptr, flags uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_vmsplice), uintptr(fd), uintptr(uiov), nrSegs, uintptr(flags), 0, 0)
	return r1, err
}

func MovePages(pid int, nrPages uintptr, pages unsafe.Pointer, nodes unsafe.Pointer, status unsafe.Pointer, flags int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_move_pages), uintptr(pid), nrPages, uintptr(pages), uintptr(nodes), uintptr(status), uintptr(flags))
	return r1, err
}

func Utimensat(dfd int, filename []byte, utimes unsafe.Pointer, flags int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_utimensat), uintptr(dfd), uintptr(unsafe.Pointer(&filename[0])), uintptr(utimes), uintptr(flags), 0, 0)
	return r1, err
}

func EpollPwait(epfd int, events unsafe.Pointer, maxevents int, timeout int, sigmask unsafe.Pointer, sigsetsize uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_epoll_pwait), uintptr(epfd), uintptr(events), uintptr(maxevents), uintptr(timeout), uintptr(sigmask), sigsetsize)
	return r1, err
}

func Signalfd(ufd int, userMask unsafe.Pointer, sizemask uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_signalfd), uintptr(ufd), uintptr(userMask), sizemask)
	return r1, err
}

func TimerfdCreate(clockid int, flags int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_timerfd_create), uintptr(clockid), uintptr(flags), 0)
	return r1, err
}

func Eventfd(count uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_eventfd), uintptr(count), 0, 0)
	return r1, err
}

func Fallocate(fd int, mode int, offset int64, length int64) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_fallocate), uintptr(fd), uintptr(mode), uintptr(offset), uintptr(length), 0, 0)
	return r1, err
}

func TimerfdSettime(ufd int, flags int, utmr unsafe.Pointer, otmr unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_timerfd_settime), uintptr(ufd), uintptr(flags), uintptr(utmr), uintptr(otmr), 0, 0)
	return r1, err
}

func TimerfdGettime(ufd int, otmr unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_timerfd_gettime), uintptr(ufd), uintptr(otmr), 0)
	return r1, err
}

func Accept4(fd int, upeerSockaddr unsafe.Pointer, upeerAddrlen unsafe.Pointer, flags int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_accept4), uintptr(fd), uintptr(upeerSockaddr), uintptr(upeerAddrlen), uintptr(flags), 0, 0)
	return r1, err
}

func Signalfd4(ufd int, userMask unsafe.Pointer, sizemask uintptr, flags int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_signalfd4), uintptr(ufd), uintptr(userMask), sizemask, uintptr(flags), 0, 0)
	return r1, err
}

func Eventfd2(count uint, flags int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_eventfd2), uintptr(count), uintptr(flags), 0)
	return r1, err
}

func EpollCreate1(flags int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_epoll_create1), uintptr(flags), 0, 0)
	return r1, err
}

func Dup3(oldfd uint, newfd uint, flags int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_dup3), uintptr(oldfd), uintptr(newfd), uintptr(flags))
	return r1, err
}

func Pipe2(fildes unsafe.Pointer, flags int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_pipe2), uintptr(fildes), uintptr(flags), 0)
	return r1, err
}

func InotifyInit1(flags int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_inotify_init1), uintptr(flags), 0, 0)
	return r1, err
}

func Preadv(fd uintptr, vec unsafe.Pointer, vlen uintptr, posL uintptr, posH uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_preadv), fd, uintptr(vec), vlen, posL, posH, 0)
	return r1, err
}

func Pwritev(fd uintptr, vec unsafe.Pointer, vlen uintptr, posL uintptr, posH uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_pwritev), fd, uintptr(vec), vlen, posL, posH, 0)
	return r1, err
}

func RtTgsigqueueinfo(tgid int, pid int, sig int, uinfo unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_rt_tgsigqueueinfo), uintptr(tgid), uintptr(pid), uintptr(sig), uintptr(uinfo), 0, 0)
	return r1, err
}

func PerfEventOpen(attrUptr unsafe.Pointer, pid int, cpu int, groupFd int, flags uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_perf_event_open), uintptr(attrUptr), uintptr(pid), uintptr(cpu), uintptr(groupFd), flags, 0)
	return r1, err
}

func Recvmmsg(fd int, mmsg unsafe.Pointer, vlen uint, flags uint, timeout unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_recvmmsg), uintptr(fd), uintptr(mmsg), uintptr(vlen), uintptr(flags), uintptr(timeout), 0)
	return r1, err
}

func FanotifyInit(flags uint, eventFFlags uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_fanotify_init), uintptr(flags), uintptr(eventFFlags), 0)
	return r1, err
}

func FanotifyMark(fanotifyFd int, flags uint, mask uint64, dfd int, pathname []byte) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_fanotify_mark), uintptr(fanotifyFd), uintptr(flags), uintptr(mask), uintptr(dfd), uintptr(unsafe.Pointer(&pathname[0])), 0)
	return r1, err
}

func Prlimit64(pid int, resource uint, newRlim unsafe.Pointer, oldRlim unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_prlimit64), uintptr(pid), uintptr(resource), uintptr(newRlim), uintptr(oldRlim), 0, 0)
	return r1, err
}

func NameToHandleAt(dfd int, name []byte, handle unsafe.Pointer, mntId unsafe.Pointer, flag int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_name_to_handle_at), uintptr(dfd), uintptr(unsafe.Pointer(&name[0])), uintptr(handle), uintptr(mntId), uintptr(flag), 0)
	return r1, err
}

func OpenByHandleAt(mountdirfd int, handle unsafe.Pointer, flags int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_open_by_handle_at), uintptr(mountdirfd), uintptr(handle), uintptr(flags))
	return r1, err
}

func ClockAdjtime(whichClock int, utx unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_clock_adjtime), uintptr(whichClock), uintptr(utx), 0)
	return r1, err
}

func Syncfs(fd int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_syncfs), uintptr(fd), 0, 0)
	return r1, err
}

func Sendmmsg(fd int, mmsg unsafe.Pointer, vlen uint, flags uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_sendmmsg), uintptr(fd), uintptr(mmsg), uintptr(vlen), uintptr(flags), 0, 0)
	return r1, err
}

func Setns(fd int, flags int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_setns), uintptr(fd), uintptr(flags), 0)
	return r1, err
}

func Getcpu(cpup unsafe.Pointer, nodep unsafe.Pointer, unused unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_getcpu), uintptr(cpup), uintptr(nodep), uintptr(unused))
	return r1, err
}

func ProcessVmReadv(pid int, lvec unsafe.Pointer, liovcnt uintptr, rvec unsafe.Pointer, riovcnt uintptr, flags uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_process_vm_readv), uintptr(pid), uintptr(lvec), liovcnt, uintptr(rvec), riovcnt, flags)
	return r1, err
}

func ProcessVmWritev(pid int, lvec unsafe.Pointer, liovcnt uintptr, rvec unsafe.Pointer, riovcnt uintptr, flags uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_process_vm_writev), uintptr(pid), uintptr(lvec), liovcnt, uintptr(rvec), riovcnt, flags)
	return r1, err
}

func Kcmp(pid1 int, pid2 int, typ int, idx1 uintptr, idx2 uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_kcmp), uintptr(pid1), uintptr(pid2), uintptr(typ), idx1, idx2, 0)
	return r1, err
}

func FinitModule(fd int, uargs []byte, flags int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_finit_module), uintptr(fd), uintptr(unsafe.Pointer(&uargs[0])), uintptr(flags))
	return r1, err
}

func SchedSetattr(pid int, uattr unsafe.Pointer, flags uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_sched_setattr), uintptr(pid), uintptr(uattr), uintptr(flags))
	return r1, err
}

func SchedGetattr(pid int, uattr unsafe.Pointer, usize uint, flags uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_sched_getattr), uintptr(pid), uintptr(uattr), uintptr(usize), uintptr(flags), 0, 0)
	return r1, err
}

func Renameat2(olddfd int, oldname []byte, newdfd int, newname []byte, flags uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_renameat2), uintptr(olddfd), uintptr(unsafe.Pointer(&oldname[0])), uintptr(newdfd), uintptr(unsafe.Pointer(&newname[0])), uintptr(flags), 0)
	return r1, err
}

func Seccomp(op uint, flags uint, uargs unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_seccomp), uintptr(op), uintptr(flags), uintptr(uargs))
	return r1, err
}

func Getrandom(ubuf unsafe.Pointer, length uintptr, flags uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_getrandom), uintptr(ubuf), length, uintptr(flags))
	return r1, err
}

func MemfdCreate(uname []byte, flags uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_memfd_create), uintptr(unsafe.Pointer(&uname[0])), uintptr(flags), 0)
	return r1, err
}

func KexecFileLoad(kernelFd int, initrdFd int, cmdlineLen uintptr, cmdlinePtr []byte, flags uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_kexec_file_load), uintptr(kernelFd), uintptr(initrdFd), cmdlineLen, uintptr(unsafe.Pointer(&cmdlinePtr[0])), flags, 0)
	return r1, err
}

func Bpf(cmd int, uattr unsafe.Pointer, size uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_bpf), uintptr(cmd), uintptr(uattr), uintptr(size))
	return r1, err
}

func Execveat(fd int, filename []byte, argv unsafe.Pointer, envp unsafe.Pointer, flags int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_execveat), uintptr(fd), uintptr(unsafe.Pointer(&filename[0])), uintptr(argv), uintptr(envp), uintptr(flags), 0)
	return r1, err
}

func Userfaultfd(flags int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_userfaultfd), uintptr(flags), 0, 0)
	return r1, err
}

func Membarrier(cmd int, flags uint, cpuId int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_membarrier), uintptr(cmd), uintptr(flags), uintptr(cpuId))
	return r1, err
}

func Mlock2(start uintptr, length uintptr, flags int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_mlock2), start, length, uintptr(flags))
	return r1, err
}

func CopyFileRange(fdIn int, offIn unsafe.Pointer, fdOut int, offOut unsafe.Pointer, length uintptr, flags uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_copy_file_range), uintptr(fdIn), uintptr(offIn), uintptr(fdOut), uintptr(offOut), length, uintptr(flags))
	return r1, err
}

func Preadv2(fd uintptr, vec unsafe.Pointer, vlen uintptr, posL uintptr, posH uintptr, flags int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_preadv2), fd, uintptr(vec), vlen, posL, posH, uintptr(flags))
	return r1, err
}

func Pwritev2(fd uintptr, vec unsafe.Pointer, vlen uintptr, posL uintptr, posH uintptr, flags int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_pwritev2), fd, uintptr(vec), vlen, posL, posH, uintptr(flags))
	return r1, err
}

func PkeyMprotect(start uintptr, length uintptr, prot uintptr, pkey int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_pkey_mprotect), start, length, prot, uintptr(pkey), 0, 0)
	return r1, err
}

func PkeyAlloc(flags uintptr, initVal uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_pkey_alloc), flags, initVal, 0)
	return r1, err
}

func PkeyFree(pkey int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_pkey_free), uintptr(pkey), 0, 0)
	return r1, err
}

func Statx(dfd int, filename []byte, flags uint, mask uint, buffer unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_statx), uintptr(dfd), uintptr(unsafe.Pointer(&filename[0])), uintptr(flags), uintptr(mask), uintptr(buffer), 0)
	return r1, err
}

func IoPgetevents(ctxId uintptr, minNr int64, nr int64, events unsafe.Pointer, timeout unsafe.Pointer, usig unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_io_pgetevents), ctxId, uintptr(minNr), uintptr(nr), uintptr(events), uintptr(timeout), uintptr(usig))
	return r1, err
}

func Rseq(rseq unsafe.Pointer, rseqLen uint32, flags int, sig uint32) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_rseq), uintptr(rseq), uintptr(rseqLen), uintptr(flags), uintptr(sig), 0, 0)
	return r1, err
}

func Uretprobe() (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_uretprobe), 0, 0, 0)
	return r1, err
}

func PidfdSendSignal(pidfd int, sig int, info unsafe.Pointer, flags uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_pidfd_send_signal), uintptr(pidfd), uintptr(sig), uintptr(info), uintptr(flags), 0, 0)
	return r1, err
}

func IoUringSetup(entries uint32, params unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_io_uring_setup), uintptr(entries), uintptr(params), 0)
	return r1, err
}

func IoUringEnter(fd uint, toSubmit uint32, minComplete uint32, flags uint32, argp unsafe.Pointer, argsz uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_io_uring_enter), uintptr(fd), uintptr(toSubmit), uintptr(minComplete), uintptr(flags), uintptr(argp), argsz)
	return r1, err
}

func IoUringRegister(fd uint, opcode uint, arg unsafe.Pointer, nrArgs uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_io_uring_register), uintptr(fd), uintptr(opcode), uintptr(arg), uintptr(nrArgs), 0, 0)
	return r1, err
}

func OpenTree(dfd int, filename []byte, flags uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_open_tree), uintptr(dfd), uintptr(unsafe.Pointer(&filename[0])), uintptr(flags))
	return r1, err
}

func MoveMount(fromDfd int, fromPathname []byte, toDfd int, toPathname []byte, flags uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_move_mount), uintptr(fromDfd), uintptr(unsafe.Pointer(&fromPathname[0])), uintptr(toDfd), uintptr(unsafe.Pointer(&toPathname[0])), uintptr(flags), 0)
	return r1, err
}

func Fsopen(_fsName []byte, flags uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_fsopen), uintptr(unsafe.Pointer(&_fsName[0])), uintptr(flags), 0)
	return r1, err
}

func Fsconfig(fd int, cmd uint, _key []byte, _value unsafe.Pointer, aux int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_fsconfig), uintptr(fd), uintptr(cmd), uintptr(unsafe.Pointer(&_key[0])), uintptr(_value), uintptr(aux), 0)
	return r1, err
}

func Fsmount(fsFd int, flags uint, attrFlags uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_fsmount), uintptr(fsFd), uintptr(flags), uintptr(attrFlags))
	return r1, err
}

func Fspick(dfd int, path []byte, flags uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_fspick), uintptr(dfd), uintptr(unsafe.Pointer(&path[0])), uintptr(flags))
	return r1, err
}

func PidfdOpen(pid int, flags uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_pidfd_open), uintptr(pid), uintptr(flags), 0)
	return r1, err
}

func Clone3(uargs unsafe.Pointer, size uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_clone3), uintptr(uargs), size, 0)
	return r1, err
}

func CloseRange(fd uint, maxFd uint, flags uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_close_range), uintptr(fd), uintptr(maxFd), uintptr(flags))
	return r1, err
}

func Openat2(dfd int, filename []byte, how unsafe.Pointer, usize uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_openat2), uintptr(dfd), uintptr(unsafe.Pointer(&filename[0])), uintptr(how), usize, 0, 0)
	return r1, err
}

func PidfdGetfd(pidfd int, fd int, flags uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_pidfd_getfd), uintptr(pidfd), uintptr(fd), uintptr(flags))
	return r1, err
}

func Faccessat2(dfd int, filename []byte, mode int, flags int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_faccessat2), uintptr(dfd), uintptr(unsafe.Pointer(&filename[0])), uintptr(mode), uintptr(flags))
	return r1, err
}

func ProcessMadvise(pidfd int, vec unsafe.Pointer, vlen uintptr, behavior int, flags uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_process_madvise), uintptr(pidfd), uintptr(vec), vlen, uintptr(behavior), uintptr(flags), 0)
	return r1, err
}

func EpollPwait2(epfd int, events unsafe.Pointer, maxevents int, timeout unsafe.Pointer, sigmask unsafe.Pointer, sigsetsize uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_epoll_pwait2), uintptr(epfd), uintptr(events), uintptr(maxevents), uintptr(timeout), uintptr(sigmask), sigsetsize)
	return r1, err
}

func MountSetattr(dfd int, path []byte, flags uint, uattr unsafe.Pointer, usize uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_mount_setattr), uintptr(dfd), uintptr(unsafe.Pointer(&path[0])), uintptr(flags), uintptr(uattr), usize, 0)
	return r1, err
}

func QuotactlFd(fd uint, cmd uint, id int, addr unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_quotactl_fd), uintptr(fd), uintptr(cmd), uintptr(id), uintptr(addr), 0, 0)
	return r1, err
}

func LandlockCreateRuleset(attr unsafe.Pointer, size uintptr, flags uint32) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_landlock_create_ruleset), uintptr(attr), size, uintptr(flags))
	return r1, err
}

func LandlockAddRule(rulesetFd int, ruleType int, ruleAttr unsafe.Pointer, flags uint32) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_landlock_add_rule), uintptr(rulesetFd), uintptr(ruleType), uintptr(ruleAttr), uintptr(flags), 0, 0)
	return r1, err
}

func LandlockRestrictSelf(rulesetFd int, flags uint32) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_landlock_restrict_self), uintptr(rulesetFd), uintptr(flags), 0)
	return r1, err
}

func MemfdSecret(flags uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_memfd_secret), uintptr(flags), 0, 0)
	return r1, err
}

func ProcessMrelease(pidfd int, flags uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_process_mrelease), uintptr(pidfd), uintptr(flags), 0)
	return r1, err
}

func FutexWaitv(waiters unsafe.Pointer, nrFutexes uint, flags uint, timeout unsafe.Pointer, clockid int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_futex_waitv), uintptr(waiters), uintptr(nrFutexes), uintptr(flags), uintptr(timeout), uintptr(clockid), 0)
	return r1, err
}

func SetMempolicyHomeNode(start uintptr, length uintptr, homeNode uintptr, flags uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_set_mempolicy_home_node), start, length, homeNode, flags, 0, 0)
	return r1, err
}

func Cachestat(fd uint, cstatRange unsafe.Pointer, cstat unsafe.Pointer, flags uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_cachestat), uintptr(fd), uintptr(cstatRange), uintptr(cstat), uintptr(flags), 0, 0)
	return r1, err
}

func Fchmodat2(dfd int, filename []byte, mode uint, flags uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_fchmodat2), uintptr(dfd), uintptr(unsafe.Pointer(&filename[0])), uintptr(mode), uintptr(flags), 0, 0)
	return r1, err
}

func MapShadowStack(addr uintptr, size uintptr, flags uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_map_shadow_stack), addr, size, uintptr(flags))
	return r1, err
}

func FutexWake(uaddr unsafe.Pointer, mask uintptr, nr int, flags uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_futex_wake), uintptr(uaddr), mask, uintptr(nr), uintptr(flags), 0, 0)
	return r1, err
}

func FutexWait(uaddr unsafe.Pointer, val uintptr, mask uintptr, flags uint, timeout unsafe.Pointer, clockid int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_futex_wait), uintptr(uaddr), val, mask, uintptr(flags), uintptr(timeout), uintptr(clockid))
	return r1, err
}

func FutexRequeue(waiters unsafe.Pointer, flags int, nrWake int, nrRequeue int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_futex_requeue), uintptr(waiters), uintptr(flags), uintptr(nrWake), uintptr(nrRequeue), 0, 0)
	return r1, err
}

func Statmount(req unsafe.Pointer, buf unsafe.Pointer, bufsize uintptr, flags uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_statmount), uintptr(req), uintptr(buf), bufsize, uintptr(flags), 0, 0)
	return r1, err
}

func Listmount(req unsafe.Pointer, mntIds unsafe.Pointer, nrMntIds uintptr, flags uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_listmount), uintptr(req), uintptr(mntIds), nrMntIds, uintptr(flags), 0, 0)
	return r1, err
}

func LsmGetSelfAttr(attr uint, ctx unsafe.Pointer, size unsafe.Pointer, flags uint32) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_lsm_get_self_attr), uintptr(attr), uintptr(ctx), uintptr(size), uintptr(flags), 0, 0)
	return r1, err
}

func LsmSetSelfAttr(attr uint, ctx unsafe.Pointer, size uint32, flags uint32) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_lsm_set_self_attr), uintptr(attr), uintptr(ctx), uintptr(size), uintptr(flags), 0, 0)
	return r1, err
}

func LsmListModules(ids unsafe.Pointer, size unsafe.Pointer, flags uint32) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_lsm_list_modules), uintptr(ids), uintptr(size), uintptr(flags))
	return r1, err
}

func Mseal(start uintptr, length uintptr, flags uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_mseal), start, length, flags)
	return r1, err
}

func Setxattrat(dfd int, pathname []byte, atFlags uint, name []byte, uargs unsafe.Pointer, usize uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_setxattrat), uintptr(dfd), uintptr(unsafe.Pointer(&pathname[0])), uintptr(atFlags), uintptr(unsafe.Pointer(&name[0])), uintptr(uargs), usize)
	return r1, err
}

func Getxattrat(dfd int, pathname []byte, atFlags uint, name []byte, uargs unsafe.Pointer, usize uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_getxattrat), uintptr(dfd), uintptr(unsafe.Pointer(&pathname[0])), uintptr(atFlags), uintptr(unsafe.Pointer(&name[0])), uintptr(uargs), usize)
	return r1, err
}

func Listxattrat(dfd int, pathname []byte, atFlags uint, list unsafe.Pointer, size uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_listxattrat), uintptr(dfd), uintptr(unsafe.Pointer(&pathname[0])), uintptr(atFlags), uintptr(list), size, 0)
	return r1, err
}

func Removexattrat(dfd int, pathname []byte, atFlags uint, name []byte) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_removexattrat), uintptr(dfd), uintptr(unsafe.Pointer(&pathname[0])), uintptr(atFlags), uintptr(unsafe.Pointer(&name[0])), 0, 0)
	return r1, err
}

func OpenTreeAttr(dfd int, filename []byte, flags uint, uattr unsafe.Pointer, usize uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_open_tree_attr), uintptr(dfd), uintptr(unsafe.Pointer(&filename[0])), uintptr(flags), uintptr(uattr), usize, 0)
	return r1, err
}

func FileSetattr(dfd int, filename []byte, ufattr unsafe.Pointer, usize uintptr, atFlags uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_file_setattr), uintptr(dfd), uintptr(unsafe.Pointer(&filename[0])), uintptr(ufattr), usize, uintptr(atFlags), 0)
	return r1, err
}

// Getsid retrieves the session ID of the process specified by pid.
// It returns the session ID and any error encountered.
func Getsid(pid int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_getsid), uintptr(pid), 0, 0)
	return r1, err
}

// Capget retrieves the capabilities of the calling thread.
// It stores the capabilities in the provided header and data structures.
// It returns 0 on success and any error encountered.
func Capget(header unsafe.Pointer, dataptr unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_capget), uintptr(header), uintptr(dataptr), 0)
	return r1, err
}

// Capset sets the capabilities of the calling thread.
// It uses the provided header and data structures.
// It returns 0 on success and any error encountered.
func Capset(header unsafe.Pointer, data unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_capset), uintptr(header), uintptr(data), 0)
	return r1, err
}

// RtSigpending examines pending signals for the calling thread.
// It stores the signal set in the provided buffer.
// It returns 0 on success and any error encountered.
func RtSigpending(uset unsafe.Pointer, sigsetsize uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_rt_sigpending), uintptr(uset), sigsetsize, 0)
	return r1, err
}

// RtSigtimedwait synchronously waits for queued signals.
// It waits for signals specified in uthese and stores info in uinfo.
// It returns the signal number and any error encountered.
func RtSigtimedwait(uthese unsafe.Pointer, uinfo unsafe.Pointer, uts unsafe.Pointer, sigsetsize uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_rt_sigtimedwait), uintptr(uthese), uintptr(uinfo), uintptr(uts), sigsetsize, 0, 0)
	return r1, err
}

// RtSigqueueinfo queues a signal and data to a process.
// It sends signal sig with info uinfo to process pid.
// It returns 0 on success and any error encountered.
func RtSigqueueinfo(pid int, sig int, uinfo unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_rt_sigqueueinfo), uintptr(pid), uintptr(sig), uintptr(uinfo))
	return r1, err
}

// RtSigsuspend replaces the signal mask and suspends the thread.
// It waits for signals not in the mask unewset.
// It returns -EINTR and any error encountered.
func RtSigsuspend(unewset unsafe.Pointer, sigsetsize uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_rt_sigsuspend), uintptr(unewset), sigsetsize, 0)
	return r1, err
}

// Sigaltstack sets and/or gets the signal stack context.
// It sets the stack to uss and gets the old stack in uoss.
// It returns 0 on success and any error encountered.
func Sigaltstack(uss unsafe.Pointer, uoss unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_sigaltstack), uintptr(uss), uintptr(uoss), 0)
	return r1, err
}

// Utime changes file timestamps.
// It sets the access and modification times of filename to times.
// It returns 0 on success and any error encountered.
func Utime(filename []byte, times unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_utime), uintptr(unsafe.Pointer(&filename[0])), uintptr(times), 0)
	return r1, err
}

// Mknod creates a special or ordinary file.
// It creates a file with name filename, mode mode, and device dev.
// It returns 0 on success and any error encountered.
func Mknod(filename []byte, mode uint, dev uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_mknod), uintptr(unsafe.Pointer(&filename[0])), uintptr(mode), uintptr(dev))
	return r1, err
}

// Personality sets the process execution domain.
// It sets the personality to personality.
// It returns the previous personality and any error encountered.
func Personality(personality uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_personality), uintptr(personality), 0, 0)
	return r1, err
}

// Ustat gets filesystem statistics.
// It retrieves statistics for device dev into ubuf.
// It returns 0 on success and any error encountered.
func Ustat(dev uint, ubuf unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_ustat), uintptr(dev), uintptr(ubuf), 0)
	return r1, err
}

// Statfs gets filesystem statistics.
// It retrieves statistics for pathname into buf.
// It returns 0 on success and any error encountered.
func Statfs(pathname []byte, buf unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_statfs), uintptr(unsafe.Pointer(&pathname[0])), uintptr(buf), 0)
	return r1, err
}

// Fstatfs gets filesystem statistics.
// It retrieves statistics for file descriptor fd into buf.
// It returns 0 on success and any error encountered.
func Fstatfs(fd uint, buf unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_fstatfs), uintptr(fd), uintptr(buf), 0)
	return r1, err
}

// Sysfs provides filesystem sysfs operations.
// It performs operation option with arguments arg1 and arg2.
// It returns the result and any error encountered.
func Sysfs(option int, arg1 uintptr, arg2 uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_sysfs), uintptr(option), arg1, arg2)
	return r1, err
}

// Getpriority gets the scheduling priority of a process.
// It gets the priority for which and who.
// It returns the priority and any error encountered.
func Getpriority(which int, who int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_getpriority), uintptr(which), uintptr(who), 0)
	return r1, err
}

// Setpriority sets the scheduling priority of a process.
// It sets the priority for which and who to niceval.
// It returns 0 on success and any error encountered.
func Setpriority(which int, who int, niceval int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_setpriority), uintptr(which), uintptr(who), uintptr(niceval))
	return r1, err
}

// SchedSetparam sets scheduling parameters.
// It sets parameters for process pid to param.
// It returns 0 on success and any error encountered.
func SchedSetparam(pid int, param unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_sched_setparam), uintptr(pid), uintptr(param), 0)
	return r1, err
}

// SchedGetparam gets scheduling parameters.
// It retrieves parameters for process pid into param.
// It returns 0 on success and any error encountered.
func SchedGetparam(pid int, param unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_sched_getparam), uintptr(pid), uintptr(param), 0)
	return r1, err
}

// SchedSetscheduler sets scheduling policy and parameters.
// It sets policy and param for process pid.
// It returns the previous policy and any error encountered.
func SchedSetscheduler(pid int, policy int, param unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_sched_setscheduler), uintptr(pid), uintptr(policy), uintptr(param))
	return r1, err
}

// SchedGetscheduler gets scheduling policy.
// It gets the policy for process pid.
// It returns the policy and any error encountered.
func SchedGetscheduler(pid int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_sched_getscheduler), uintptr(pid), 0, 0)
	return r1, err
}

// SchedGetPriorityMax gets maximum priority for scheduling policy.
// It gets the max priority for policy.
// It returns the max priority and any error encountered.
func SchedGetPriorityMax(policy int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_sched_get_priority_max), uintptr(policy), 0, 0)
	return r1, err
}

// SchedGetPriorityMin gets minimum priority for scheduling policy.
// It gets the min priority for policy.
// It returns the min priority and any error encountered.
func SchedGetPriorityMin(policy int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_sched_get_priority_min), uintptr(policy), 0, 0)
	return r1, err
}

// SchedRrGetInterval gets the SCHED_RR interval for a process.
// It gets the interval for process pid into interval.
// It returns 0 on success and any error encountered.
func SchedRrGetInterval(pid int, interval unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_sched_rr_get_interval), uintptr(pid), uintptr(interval), 0)
	return r1, err
}

// Mlock locks pages in memory.
// It locks len bytes starting at start.
// It returns 0 on success and any error encountered.
func Mlock(start uintptr, length uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_mlock), start, length, 0)
	return r1, err
}

// Munlock unlocks pages in memory.
// It unlocks len bytes starting at start.
// It returns 0 on success and any error encountered.
func Munlock(start uintptr, length uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_munlock), start, length, 0)
	return r1, err
}

// Mlockall locks all pages in memory.
// It locks all pages with flags.
// It returns 0 on success and any error encountered.
func Mlockall(flags int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_mlockall), uintptr(flags), 0, 0)
	return r1, err
}

// Munlockall unlocks all pages in memory.
// It unlocks all locked pages.
// It returns 0 on success and any error encountered.
func Munlockall() (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_munlockall), 0, 0, 0)
	return r1, err
}

// Vhangup virtually hangs up the current terminal.
// It sends a hangup signal to the controlling terminal.
// It returns 0 on success and any error encountered.
func Vhangup() (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_vhangup), 0, 0, 0)
	return r1, err
}

// ModifyLdt modifies the local descriptor table.
// It performs function func_ with ptr and bytecount.
// It returns the result and any error encountered.
func ModifyLdt(func_ int, ptr unsafe.Pointer, bytecount uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_modify_ldt), uintptr(func_), uintptr(ptr), uintptr(bytecount))
	return r1, err
}

// PivotRoot changes the root filesystem.
// It exchanges new_root and put_old.
// It returns 0 on success and any error encountered.
func PivotRoot(newRoot []byte, putOld []byte) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_pivot_root), uintptr(unsafe.Pointer(&newRoot[0])), uintptr(unsafe.Pointer(&putOld[0])), 0)
	return r1, err
}

// Prctl performs process control operations.
// It performs option with arguments arg2, arg3, arg4, arg5.
// It returns the result and any error encountered.
func Prctl(option int, arg2 uintptr, arg3 uintptr, arg4 uintptr, arg5 uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_prctl), uintptr(option), arg2, arg3, arg4, arg5, 0)
	return r1, err
}

// ArchPrctl performs architecture-specific process control.
// It performs option with arg2.
// It returns 0 on success and any error encountered.
func ArchPrctl(option int, arg2 uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_arch_prctl), uintptr(option), arg2, 0)
	return r1, err
}

// Adjtimex tunes kernel time variables.
// It adjusts time parameters in txc_p.
// It returns the clock state and any error encountered.
func Adjtimex(txc_p unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_adjtimex), uintptr(txc_p), 0, 0)
	return r1, err
}

// Setrlimit sets resource limits.
// It sets the limit for resource to rlim.
// It returns 0 on success and any error encountered.
func Setrlimit(resource uint, rlim unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_setrlimit), uintptr(resource), uintptr(rlim), 0)
	return r1, err
}

// Chroot changes the root directory.
// It changes the root to filename.
// It returns 0 on success and any error encountered.
func Chroot(filename []byte) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_chroot), uintptr(unsafe.Pointer(&filename[0])), 0, 0)
	return r1, err
}

// Sync synchronizes filesystem data.
// It flushes all pending disk writes.
// It returns 0 on success and any error encountered.
func Sync() (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_sync), 0, 0, 0)
	return r1, err
}

// Acct enables or disables process accounting.
// It enables accounting to name.
// It returns 0 on success and any error encountered.
func Acct(name []byte) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_acct), uintptr(unsafe.Pointer(&name[0])), 0, 0)
	return r1, err
}

// Settimeofday sets the system time and timezone.
// It sets time to tv and timezone to tz.
// It returns 0 on success and any error encountered.
func Settimeofday(tv unsafe.Pointer, tz unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_settimeofday), uintptr(tv), uintptr(tz), 0)
	return r1, err
}

// Mount mounts a filesystem.
// It mounts dev_name at dir_name with type and flags, using data.
// It returns 0 on success and any error encountered.
func Mount(devName []byte, dirName []byte, typ []byte, flags uintptr, data unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_mount), uintptr(unsafe.Pointer(&devName[0])), uintptr(unsafe.Pointer(&dirName[0])), uintptr(unsafe.Pointer(&typ[0])), flags, uintptr(data), 0)
	return r1, err
}

// Umount unmounts a filesystem.
// It unmounts name with flags.
// It returns 0 on success and any error encountered.
func Umount(name []byte, flags int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_umount), uintptr(unsafe.Pointer(&name[0])), uintptr(flags), 0)
	return r1, err
}

// Swapon starts swapping to a file.
// It starts swapping to specialfile with swap_flags.
// It returns 0 on success and any error encountered.
func Swapon(specialfile []byte, swapFlags int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_swapon), uintptr(unsafe.Pointer(&specialfile[0])), uintptr(swapFlags), 0)
	return r1, err
}

// Swapoff stops swapping to a file.
// It stops swapping to specialfile.
// It returns 0 on success and any error encountered.
func Swapoff(specialfile []byte) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_swapoff), uintptr(unsafe.Pointer(&specialfile[0])), 0, 0)
	return r1, err
}

// Reboot reboots or halts the system.
// It performs cmd with magic1 and magic2, using arg.
// It returns 0 on success and any error encountered.
func Reboot(magic1 int, magic2 int, cmd uint, arg unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_reboot), uintptr(magic1), uintptr(magic2), uintptr(cmd), uintptr(arg), 0, 0)
	return r1, err
}

// Sethostname sets the hostname.
// It sets the hostname to name of length len.
// It returns 0 on success and any error encountered.
func Sethostname(name []byte, length int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_sethostname), uintptr(unsafe.Pointer(&name[0])), uintptr(length), 0)
	return r1, err
}

// Setdomainname sets the domain name.
// It sets the domain name to name of length len.
// It returns 0 on success and any error encountered.
func Setdomainname(name []byte, length int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_setdomainname), uintptr(unsafe.Pointer(&name[0])), uintptr(length), 0)
	return r1, err
}

// Iopl sets the I/O privilege level.
// It sets the I/O privilege level to level.
// It returns 0 on success and any error encountered.
func Iopl(level uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_iopl), uintptr(level), 0, 0)
	return r1, err
}

// Ioperm sets port access permissions.
// It sets permissions for ports from to num, turning on if turn_on.
// It returns 0 on success and any error encountered.
func Ioperm(from uintptr, num uint, turnOn int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_ioperm), from, uintptr(num), uintptr(turnOn))
	return r1, err
}

// InitModule loads a kernel module.
// It loads module umod of length len with args uargs.
// It returns 0 on success and any error encountered.
func InitModule(umod unsafe.Pointer, length uintptr, uargs []byte) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_init_module), uintptr(umod), length, uintptr(unsafe.Pointer(&uargs[0])))
	return r1, err
}

// DeleteModule unloads a kernel module.
// It unloads module name_user with flags.
// It returns 0 on success and any error encountered.
func DeleteModule(nameUser []byte, flags uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_delete_module), uintptr(unsafe.Pointer(&nameUser[0])), uintptr(flags), 0)
	return r1, err
}

// Quotactl manipulates disk quotas.
// It performs cmd on special for id with addr.
// It returns 0 on success and any error encountered.
func Quotactl(cmd uint, special []byte, id int, addr unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_quotactl), uintptr(cmd), uintptr(unsafe.Pointer(&special[0])), uintptr(id), uintptr(addr), 0, 0)
	return r1, err
}

// Gettid gets the thread ID.
// It returns the thread ID of the calling thread.
func Gettid() (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_gettid), 0, 0, 0)
	return r1, err
}

// Readahead initiates readahead.
// It starts readahead for fd at offset for count bytes.
// It returns 0 on success and any error encountered.
func Readahead(fd int, offset int64, count uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_readahead), uintptr(fd), uintptr(offset), count)
	return r1, err
}

// Setxattr sets an extended attribute.
// It sets attribute name to value on pathname with flags.
// It returns 0 on success and any error encountered.
func Setxattr(pathname []byte, name []byte, value unsafe.Pointer, size uintptr, flags int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_setxattr), uintptr(unsafe.Pointer(&pathname[0])), uintptr(unsafe.Pointer(&name[0])), uintptr(value), size, uintptr(flags), 0)
	return r1, err
}

// Lsetxattr sets an extended attribute on a symbolic link.
// It sets attribute name to value on pathname with flags.
// It returns 0 on success and any error encountered.
func Lsetxattr(pathname []byte, name []byte, value unsafe.Pointer, size uintptr, flags int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_lsetxattr), uintptr(unsafe.Pointer(&pathname[0])), uintptr(unsafe.Pointer(&name[0])), uintptr(value), size, uintptr(flags), 0)
	return r1, err
}

// Fsetxattr sets an extended attribute on a file descriptor.
// It sets attribute name to value on fd with flags.
// It returns 0 on success and any error encountered.
func Fsetxattr(fd int, name []byte, value unsafe.Pointer, size uintptr, flags int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_fsetxattr), uintptr(fd), uintptr(unsafe.Pointer(&name[0])), uintptr(value), size, uintptr(flags), 0)
	return r1, err
}

// Getxattr gets an extended attribute.
// It retrieves attribute name from pathname into value.
// It returns the size and any error encountered.
func Getxattr(pathname []byte, name []byte, value unsafe.Pointer, size uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_getxattr), uintptr(unsafe.Pointer(&pathname[0])), uintptr(unsafe.Pointer(&name[0])), uintptr(value), size, 0, 0)
	return r1, err
}

// Lgetxattr gets an extended attribute from a symbolic link.
// It retrieves attribute name from pathname into value.
// It returns the size and any error encountered.
func Lgetxattr(pathname []byte, name []byte, value unsafe.Pointer, size uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_lgetxattr), uintptr(unsafe.Pointer(&pathname[0])), uintptr(unsafe.Pointer(&name[0])), uintptr(value), size, 0, 0)
	return r1, err
}

// Fgetxattr gets an extended attribute from a file descriptor.
// It retrieves attribute name from fd into value.
// It returns the size and any error encountered.
func Fgetxattr(fd int, name []byte, value unsafe.Pointer, size uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_fgetxattr), uintptr(fd), uintptr(unsafe.Pointer(&name[0])), uintptr(value), size, 0, 0)
	return r1, err
}

// Listxattr lists extended attributes.
// It lists attributes of pathname into list.
// It returns the size and any error encountered.
func Listxattr(pathname []byte, list unsafe.Pointer, size uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_listxattr), uintptr(unsafe.Pointer(&pathname[0])), uintptr(list), size)
	return r1, err
}

// Llistxattr lists extended attributes of a symbolic link.
// It lists attributes of pathname into list.
// It returns the size and any error encountered.
func Llistxattr(pathname []byte, list unsafe.Pointer, size uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_llistxattr), uintptr(unsafe.Pointer(&pathname[0])), uintptr(list), size)
	return r1, err
}

// Flistxattr lists extended attributes of a file descriptor.
// It lists attributes of fd into list.
// It returns the size and any error encountered.
func Flistxattr(fd int, list unsafe.Pointer, size uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_flistxattr), uintptr(fd), uintptr(list), size)
	return r1, err
}

// Removexattr removes an extended attribute.
// It removes attribute name from pathname.
// It returns 0 on success and any error encountered.
func Removexattr(pathname []byte, name []byte) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_removexattr), uintptr(unsafe.Pointer(&pathname[0])), uintptr(unsafe.Pointer(&name[0])), 0)
	return r1, err
}

// Lremovexattr removes an extended attribute from a symbolic link.
// It removes attribute name from pathname.
// It returns 0 on success and any error encountered.
func Lremovexattr(pathname []byte, name []byte) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_lremovexattr), uintptr(unsafe.Pointer(&pathname[0])), uintptr(unsafe.Pointer(&name[0])), 0)
	return r1, err
}

// Fremovexattr removes an extended attribute from a file descriptor.
// It removes attribute name from fd.
// It returns 0 on success and any error encountered.
func Fremovexattr(fd int, name []byte) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_fremovexattr), uintptr(fd), uintptr(unsafe.Pointer(&name[0])), 0)
	return r1, err
}

// Tkill sends a signal to a thread.
// It sends signal sig to thread pid.
// It returns 0 on success and any error encountered.
func Tkill(pid int, sig int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_tkill), uintptr(pid), uintptr(sig), 0)
	return r1, err
}

// Time gets the current time.
// It stores the time in tloc.
// It returns the time and any error encountered.
func Time(tloc unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_time), uintptr(tloc), 0, 0)
	return r1, err
}

// Futex performs a fast user-space locking operation.
// It operates on uaddr with op, val, utime, uaddr2, val3.
// It returns 0 on success and any error encountered.
func Futex(uaddr unsafe.Pointer, op int, val int, utime unsafe.Pointer, uaddr2 unsafe.Pointer, val3 int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_futex), uintptr(uaddr), uintptr(op), uintptr(val), uintptr(utime), uintptr(uaddr2), uintptr(val3))
	return r1, err
}

// SchedSetaffinity sets CPU affinity.
// It sets affinity of pid to user_mask_ptr of length len.
// It returns 0 on success and any error encountered.
func SchedSetaffinity(pid int, length uint, userMaskPtr unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_sched_setaffinity), uintptr(pid), uintptr(length), uintptr(userMaskPtr))
	return r1, err
}

// SchedGetaffinity gets CPU affinity.
// It gets affinity of pid into user_mask_ptr of length len.
// It returns 0 on success and any error encountered.
func SchedGetaffinity(pid int, length uint, userMaskPtr unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_sched_getaffinity), uintptr(pid), uintptr(length), uintptr(userMaskPtr))
	return r1, err
}

// IoSetup creates an asynchronous I/O context.
// It creates a context with nr_events into ctxp.
// It returns 0 on success and any error encountered.
func IoSetup(nrEvents uintptr, ctxp unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_io_setup), nrEvents, uintptr(ctxp), 0)
	return r1, err
}

// IoDestroy destroys an asynchronous I/O context.
// It destroys context ctx.
// It returns 0 on success and any error encountered.
func IoDestroy(ctx uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_io_destroy), ctx, 0, 0)
	return r1, err
}

// IoGetevents reads asynchronous I/O events.
// It reads up to nr events from ctx_id into events, with timeout.
// It returns the number of events and any error encountered.
func IoGetevents(ctxId uintptr, minNr int64, nr int64, events unsafe.Pointer, timeout unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_io_getevents), ctxId, uintptr(minNr), uintptr(nr), uintptr(events), uintptr(timeout), 0)
	return r1, err
}

// IoSubmit submits asynchronous I/O blocks.
// It submits nr blocks from iocbpp to ctx_id.
// It returns the number submitted and any error encountered.
func IoSubmit(ctxId uintptr, nr int64, iocbpp unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_io_submit), ctxId, uintptr(nr), uintptr(iocbpp))
	return r1, err
}

// IoCancel cancels an asynchronous I/O request.
// It cancels iocb in ctx_id into result.
// It returns 0 on success and any error encountered.
func IoCancel(ctxId uintptr, iocb unsafe.Pointer, result unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_io_cancel), ctxId, uintptr(iocb), uintptr(result))
	return r1, err
}

// EpollCreate creates an epoll instance.
// It creates an epoll instance with size.
// It returns the file descriptor and any error encountered.
func EpollCreate(size int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_epoll_create), uintptr(size), 0, 0)
	return r1, err
}

// RemapFilePages remaps file pages.
// It remaps pages at start with size, prot, pgoff, flags.
// It returns 0 on success and any error encountered.
func RemapFilePages(start uintptr, size uintptr, prot uintptr, pgoff uintptr, flags uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_remap_file_pages), start, size, prot, pgoff, flags, 0)
	return r1, err
}

// Getdents64 gets directory entries.
// It reads directory entries from fd into dirent.
// It returns the number of bytes read and any error encountered.
func Getdents64(fd uint, dirent unsafe.Pointer, count uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_getdents64), uintptr(fd), uintptr(dirent), uintptr(count))
	return r1, err
}

// SetTidAddress sets the clear_child_tid address.
// It sets the address to tidptr.
// It returns the old tid and any error encountered.
func SetTidAddress(tidptr unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_set_tid_address), uintptr(tidptr), 0, 0)
	return r1, err
}

// RestartSyscall restarts a system call after interruption.
// It restarts the interrupted system call.
// It returns the result and any error encountered.
func RestartSyscall() (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_restart_syscall), 0, 0, 0)
	return r1, err
}

// Semtimedop performs semaphore operations with timeout.
// It performs nsops operations in tsops on semid with timeout.
// It returns 0 on success and any error encountered.
func Semtimedop(semid int, tsops unsafe.Pointer, nsops uint, timeout unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_semtimedop), uintptr(semid), uintptr(tsops), uintptr(nsops), uintptr(timeout), 0, 0)
	return r1, err
}

// Fadvise64 gives advice about file access patterns.
// It gives advice about fd from offset for len bytes.
// It returns 0 on success and any error encountered.
func Fadvise64(fd int, offset int64, length int64, advice int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_fadvise64), uintptr(fd), uintptr(offset), uintptr(length), uintptr(advice), 0, 0)
	return r1, err
}

// TimerCreate creates a POSIX timer.
// It creates a timer with clock which_clock, event spec, into created_timer_id.
// It returns 0 on success and any error encountered.
func TimerCreate(whichClock int, timerEventSpec unsafe.Pointer, createdTimerId unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_timer_create), uintptr(whichClock), uintptr(timerEventSpec), uintptr(createdTimerId))
	return r1, err
}

// TimerSettime sets the time of a POSIX timer.
// It sets timer_id with flags, new_setting, into old_setting.
// It returns 0 on success and any error encountered.
func TimerSettime(timerId uintptr, flags int, newSetting unsafe.Pointer, oldSetting unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_timer_settime), timerId, uintptr(flags), uintptr(newSetting), uintptr(oldSetting), 0, 0)
	return r1, err
}

// TimerGettime gets the time of a POSIX timer.
// It gets timer_id into setting.
// It returns 0 on success and any error encountered.
func TimerGettime(timerId uintptr, setting unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_timer_gettime), timerId, uintptr(setting), 0)
	return r1, err
}

// TimerGetoverrun gets overrun count of a POSIX timer.
// It gets overrun for timer_id.
// It returns the overrun count and any error encountered.
func TimerGetoverrun(timerId uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_timer_getoverrun), timerId, 0, 0)
	return r1, err
}

// TimerDelete deletes a POSIX timer.
// It deletes timer_id.
// It returns 0 on success and any error encountered.
func TimerDelete(timerId uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_timer_delete), timerId, 0, 0)
	return r1, err
}

// ClockSettime sets the time of a clock.
// It sets clock which_clock to tp.
// It returns 0 on success and any error encountered.
func ClockSettime(whichClock int, tp unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_clock_settime), uintptr(whichClock), uintptr(tp), 0)
	return r1, err
}

// ClockGettime gets the time of a clock.
// It gets clock which_clock into tp.
// It returns 0 on success and any error encountered.
func ClockGettime(whichClock int, tp unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_clock_gettime), uintptr(whichClock), uintptr(tp), 0)
	return r1, err
}

// ClockGetres gets the resolution of a clock.
// It gets resolution of which_clock into tp.
// It returns 0 on success and any error encountered.
func ClockGetres(whichClock int, tp unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_clock_getres), uintptr(whichClock), uintptr(tp), 0)
	return r1, err
}

// ClockNanosleep sleeps for a specified time.
// It sleeps on which_clock with flags, rqtp, into rmtp.
// It returns 0 on success and any error encountered.
func ClockNanosleep(whichClock int, flags int, rqtp unsafe.Pointer, rmtp unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_clock_nanosleep), uintptr(whichClock), uintptr(flags), uintptr(rqtp), uintptr(rmtp), 0, 0)
	return r1, err
}

// ExitGroup exits all threads in the process.
// It exits with error_code.
// It does not return.
func ExitGroup(errorCode int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_exit_group), uintptr(errorCode), 0, 0)
	return r1, err
}

// EpollWait waits for events on an epoll instance.
// It waits on epfd for maxevents into events, with timeout.
// It returns the number of events and any error encountered.
func EpollWait(epfd int, events unsafe.Pointer, maxevents int, timeout int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_epoll_wait), uintptr(epfd), uintptr(events), uintptr(maxevents), uintptr(timeout), 0, 0)
	return r1, err
}

// EpollCtl controls an epoll instance.
// It performs op on fd with event for epfd.
// It returns 0 on success and any error encountered.
func EpollCtl(epfd int, op int, fd int, event unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_epoll_ctl), uintptr(epfd), uintptr(op), uintptr(fd), uintptr(event), 0, 0)
	return r1, err
}

// Tgkill sends a signal to a thread group.
// It sends sig to tgid/pid.
// It returns 0 on success and any error encountered.
func Tgkill(tgid int, pid int, sig int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_tgkill), uintptr(tgid), uintptr(pid), uintptr(sig))
	return r1, err
}

// Utimes sets file timestamps.
// It sets timestamps of filename to utimes.
// It returns 0 on success and any error encountered.
func Utimes(filename []byte, utimes unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_utimes), uintptr(unsafe.Pointer(&filename[0])), uintptr(utimes), 0)
	return r1, err
}

// Mbind sets memory policy for a memory range.
// It sets policy for start to length with mode, nmask, maxnode, flags.
// It returns 0 on success and any error encountered.
func Mbind(start uintptr, length uintptr, mode uintptr, nmask unsafe.Pointer, maxnode uintptr, flags uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_mbind), start, length, mode, uintptr(nmask), maxnode, flags)
	return r1, err
}

// SetMempolicy sets the default memory policy.
// It sets policy to mode with nmask, maxnode.
// It returns 0 on success and any error encountered.
func SetMempolicy(mode int, nmask unsafe.Pointer, maxnode uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_set_mempolicy), uintptr(mode), uintptr(nmask), maxnode)
	return r1, err
}

// GetMempolicy gets the memory policy.
// It gets policy into nmask for maxnode, addr, flags.
// It returns 0 on success and any error encountered.
func GetMempolicy(policy unsafe.Pointer, nmask unsafe.Pointer, maxnode uintptr, addr uintptr, flags uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_get_mempolicy), uintptr(policy), uintptr(nmask), maxnode, addr, flags, 0)
	return r1, err
}

// MqOpen opens a message queue.
// It opens queue u_name with oflag, mode, u_attr.
// It returns the descriptor and any error encountered.
func MqOpen(uName []byte, oflag int, mode uint, uAttr unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_mq_open), uintptr(unsafe.Pointer(&uName[0])), uintptr(oflag), uintptr(mode), uintptr(uAttr), 0, 0)
	return r1, err
}

// MqUnlink removes a message queue.
// It removes queue u_name.
// It returns 0 on success and any error encountered.
func MqUnlink(uName []byte) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_mq_unlink), uintptr(unsafe.Pointer(&uName[0])), 0, 0)
	return r1, err
}

// MqTimedsend sends a message to a message queue with timeout.
// It sends u_msg_ptr of msg_len with msg_prio to mqdes, with u_abs_timeout.
// It returns 0 on success and any error encountered.
func MqTimedsend(mqdes uintptr, uMsgPtr unsafe.Pointer, msgLen uintptr, msgPrio uint, uAbsTimeout unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_mq_timedsend), mqdes, uintptr(uMsgPtr), msgLen, uintptr(msgPrio), uintptr(uAbsTimeout), 0)
	return r1, err
}

// MqTimedreceive receives a message from a message queue with timeout.
// It receives into u_msg_ptr of msg_len, msg_prio, with u_abs_timeout from mqdes.
// It returns the size and any error encountered.
func MqTimedreceive(mqdes uintptr, uMsgPtr unsafe.Pointer, msgLen uintptr, uMsgPrio unsafe.Pointer, uAbsTimeout unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_mq_timedreceive), mqdes, uintptr(uMsgPtr), msgLen, uintptr(uMsgPrio), uintptr(uAbsTimeout), 0)
	return r1, err
}

// MqNotify registers for notification of message arrival.
// It registers u_notification for mqdes.
// It returns 0 on success and any error encountered.
func MqNotify(mqdes uintptr, uNotification unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_mq_notify), mqdes, uintptr(uNotification), 0)
	return r1, err
}

// MqGetsetattr gets and sets message queue attributes.
// It gets/sets u_mqstat for mqdes into u_omqstat.
// It returns 0 on success and any error encountered.
func MqGetsetattr(mqdes uintptr, uMqstat unsafe.Pointer, uOmqstat unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_mq_getsetattr), mqdes, uintptr(uMqstat), uintptr(uOmqstat))
	return r1, err
}

// KexecLoad loads a new kernel.
// It loads kernel at entry with nr_segments segments, flags.
// It returns 0 on success and any error encountered.
func KexecLoad(entry uintptr, nrSegments uintptr, segments unsafe.Pointer, flags uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_kexec_load), entry, nrSegments, uintptr(segments), flags, 0, 0)
	return r1, err
}

// Waitid waits for process state changes.
// It waits for which upid with options, into infop, ru.
// It returns 0 on success and any error encountered.
func Waitid(which int, upid int, infop unsafe.Pointer, options int, ru unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_waitid), uintptr(which), uintptr(upid), uintptr(infop), uintptr(options), uintptr(ru), 0)
	return r1, err
}

// AddKey adds a key to the kernel keyring.
// It adds key of _type _description with _payload of plen to ringid.
// It returns the key ID and any error encountered.
func AddKey(_type []byte, _description []byte, _payload unsafe.Pointer, plen uintptr, ringid int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_add_key), uintptr(unsafe.Pointer(&_type[0])), uintptr(unsafe.Pointer(&_description[0])), uintptr(_payload), plen, uintptr(ringid), 0)
	return r1, err
}

// RequestKey requests a key from the kernel.
// It requests key of _type _description with _callout_info to destringid.
// It returns the key ID and any error encountered.
func RequestKey(_type []byte, _description []byte, _calloutInfo []byte, destringid int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_request_key), uintptr(unsafe.Pointer(&_type[0])), uintptr(unsafe.Pointer(&_description[0])), uintptr(unsafe.Pointer(&_calloutInfo[0])), uintptr(destringid), 0, 0)
	return r1, err
}

// Keyctl performs key management operations.
// It performs option with arg2, arg3, arg4, arg5.
// It returns the result and any error encountered.
func Keyctl(option int, arg2 uintptr, arg3 uintptr, arg4 uintptr, arg5 uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_keyctl), uintptr(option), arg2, arg3, arg4, arg5, 0)
	return r1, err
}

// IoprioSet sets I/O scheduling priority.
// It sets priority for which who to ioprio.
// It returns 0 on success and any error encountered.
func IoprioSet(which int, who int, ioprio int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_ioprio_set), uintptr(which), uintptr(who), uintptr(ioprio))
	return r1, err
}

// IoprioGet gets I/O scheduling priority.
// It gets priority for which who.
// It returns the priority and any error encountered.
func IoprioGet(which int, who int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_ioprio_get), uintptr(which), uintptr(who), 0)
	return r1, err
}

// InotifyInit initializes an inotify instance.
// It creates an inotify file descriptor.
// It returns the descriptor and any error encountered.
func InotifyInit() (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_inotify_init), 0, 0, 0)
	return r1, err
}

// InotifyAddWatch adds a watch to an inotify instance.
// It adds watch for pathname with mask to fd.
// It returns the watch descriptor and any error encountered.
func InotifyAddWatch(fd int, pathname []byte, mask uint32) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_inotify_add_watch), uintptr(fd), uintptr(unsafe.Pointer(&pathname[0])), uintptr(mask))
	return r1, err
}

// InotifyRmWatch removes a watch from an inotify instance.
// It removes watch wd from fd.
// It returns 0 on success and any error encountered.
func InotifyRmWatch(fd int, wd int32) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_inotify_rm_watch), uintptr(fd), uintptr(wd), 0)
	return r1, err
}

// MigratePages migrates pages to other nodes.
// It migrates pages of pid with maxnode, old_nodes, new_nodes.
// It returns the number of pages not migrated and any error encountered.
func MigratePages(pid int, maxnode uintptr, oldNodes unsafe.Pointer, newNodes unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_migrate_pages), uintptr(pid), maxnode, uintptr(oldNodes), uintptr(newNodes), 0, 0)
	return r1, err
}

// Openat opens a file relative to a directory.
// It opens filename with flags, mode relative to dfd.
// It returns the descriptor and any error encountered.
func Openat(dfd int, filename []byte, flags int, mode uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_openat), uintptr(dfd), uintptr(unsafe.Pointer(&filename[0])), uintptr(flags), uintptr(mode), 0, 0)
	return r1, err
}

// Mkdirat creates a directory relative to a directory.
// It creates pathname with mode relative to dfd.
// It returns 0 on success and any error encountered.
func Mkdirat(dfd int, pathname []byte, mode uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_mkdirat), uintptr(dfd), uintptr(unsafe.Pointer(&pathname[0])), uintptr(mode))
	return r1, err
}

// Mknodat creates a special file relative to a directory.
// It creates filename with mode, dev relative to dfd.
// It returns 0 on success and any error encountered.
func Mknodat(dfd int, filename []byte, mode uint, dev uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_mknodat), uintptr(dfd), uintptr(unsafe.Pointer(&filename[0])), uintptr(mode), uintptr(dev), 0, 0)
	return r1, err
}

// Fchownat changes ownership of a file relative to a directory.
// It changes user, group of filename relative to dfd with flag.
// It returns 0 on success and any error encountered.
func Fchownat(dfd int, filename []byte, user uint, group uint, flag int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_fchownat), uintptr(dfd), uintptr(unsafe.Pointer(&filename[0])), uintptr(user), uintptr(group), uintptr(flag), 0)
	return r1, err
}

// Futimesat sets timestamps of a file relative to a directory.
// It sets utimes for filename relative to dfd.
// It returns 0 on success and any error encountered.
func Futimesat(dfd int, filename []byte, utimes unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_futimesat), uintptr(dfd), uintptr(unsafe.Pointer(&filename[0])), uintptr(utimes))
	return r1, err
}

// Newfstatat gets file status relative to a directory.
// It gets stat of filename relative to dfd with flag into statbuf.
// It returns 0 on success and any error encountered.
func Newfstatat(dfd int, filename []byte, statbuf unsafe.Pointer, flag int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_newfstatat), uintptr(dfd), uintptr(unsafe.Pointer(&filename[0])), uintptr(statbuf), uintptr(flag), 0, 0)
	return r1, err
}

// Unlinkat removes a file relative to a directory.
// It removes pathname relative to dfd with flag.
// It returns 0 on success and any error encountered.
func Unlinkat(dfd int, pathname []byte, flag int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_unlinkat), uintptr(dfd), uintptr(unsafe.Pointer(&pathname[0])), uintptr(flag))
	return r1, err
}

// Renameat renames a file relative to directories.
// It renames oldname to newname relative to olddfd/newdfd.
// It returns 0 on success and any error encountered.
func Renameat(olddfd int, oldname []byte, newdfd int, newname []byte) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_renameat), uintptr(olddfd), uintptr(unsafe.Pointer(&oldname[0])), uintptr(newdfd), uintptr(unsafe.Pointer(&newname[0])), 0, 0)
	return r1, err
}

// Linkat creates a link relative to directories.
// It links oldname to newname relative to olddfd/newdfd with flags.
// It returns 0 on success and any error encountered.
func Linkat(olddfd int, oldname []byte, newdfd int, newname []byte, flags int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_linkat), uintptr(olddfd), uintptr(unsafe.Pointer(&oldname[0])), uintptr(newdfd), uintptr(unsafe.Pointer(&newname[0])), uintptr(flags), 0)
	return r1, err
}

// Symlinkat creates a symbolic link relative to a directory.
// It creates link from oldname to newname relative to newdfd.
// It returns 0 on success and any error encountered.
func Symlinkat(oldname []byte, newdfd int, newname []byte) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_symlinkat), uintptr(unsafe.Pointer(&oldname[0])), uintptr(newdfd), uintptr(unsafe.Pointer(&newname[0])))
	return r1, err
}

// Readlinkat reads the contents of a symbolic link relative to a directory.
// It reads link pathname relative to dfd into buf of bufsiz.
// It returns the number of bytes read and any error encountered.
func Readlinkat(dfd int, pathname []byte, buf unsafe.Pointer, bufsiz int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_readlinkat), uintptr(dfd), uintptr(unsafe.Pointer(&pathname[0])), uintptr(buf), uintptr(bufsiz), 0, 0)
	return r1, err
}

// Fchmodat changes permissions of a file relative to a directory.
// It changes mode of filename relative to dfd.
// It returns 0 on success and any error encountered.
func Fchmodat(dfd int, filename []byte, mode uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_fchmodat), uintptr(dfd), uintptr(unsafe.Pointer(&filename[0])), uintptr(mode))
	return r1, err
}

// Faccessat checks access to a file relative to a directory.
// It checks mode access to filename relative to dfd.
// It returns 0 on success and any error encountered.
func Faccessat(dfd int, filename []byte, mode int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_faccessat), uintptr(dfd), uintptr(unsafe.Pointer(&filename[0])), uintptr(mode))
	return r1, err
}

// Pselect6 is a variant of select with signal mask.
// It selects with n, inp, outp, exp, tsp, sig.
// It returns the number of ready descriptors and any error encountered.
func Pselect6(n int, inp unsafe.Pointer, outp unsafe.Pointer, exp unsafe.Pointer, tsp unsafe.Pointer, sig unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_pselect6), uintptr(n), uintptr(inp), uintptr(outp), uintptr(exp), uintptr(tsp), uintptr(sig))
	return r1, err
}

// Ppoll polls file descriptors with signal mask.
// It polls ufds with nfds, tsp, sigmask, sigsetsize.
// It returns the number of ready descriptors and any error encountered.
func Ppoll(ufds unsafe.Pointer, nfds uintptr, tsp unsafe.Pointer, sigmask unsafe.Pointer, sigsetsize uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_ppoll), uintptr(ufds), nfds, uintptr(tsp), uintptr(sigmask), sigsetsize, 0)
	return r1, err
}

// Unshare disassociates parts of the process execution context.
// It unshares unshare_flags.
// It returns 0 on success and any error encountered.
func Unshare(unshareFlags uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_unshare), unshareFlags, 0, 0)
	return r1, err
}

// SetRobustList sets the robust futex list.
// It sets head with length.
// It returns 0 on success and any error encountered.
func SetRobustList(head unsafe.Pointer, length uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_set_robust_list), uintptr(head), length, 0)
	return r1, err
}

// GetRobustList gets the robust futex list.
// It gets head and len_ptr for pid.
// It returns 0 on success and any error encountered.
func GetRobustList(pid int, headPtr unsafe.Pointer, lenPtr unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_get_robust_list), uintptr(pid), uintptr(headPtr), uintptr(lenPtr))
	return r1, err
}

// Splice splices data between files.
// It splices len bytes from fd_in off_in to fd_out off_out with flags.
// It returns the number of bytes spliced and any error encountered.
func Splice(fdIn int, offIn unsafe.Pointer, fdOut int, offOut unsafe.Pointer, length uintptr, flags uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_splice), uintptr(fdIn), uintptr(offIn), uintptr(fdOut), uintptr(offOut), length, uintptr(flags))
	return r1, err
}

// Tee duplicates data between pipes.
// It tees len bytes from fdin to fdout with flags.
// It returns the number of bytes teed and any error encountered.
func Tee(fdin int, fdout int, length uintptr, flags uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_tee), uintptr(fdin), uintptr(fdout), length, uintptr(flags), 0, 0)
	return r1, err
}

// SyncFileRange syncs a file segment.
// It syncs fd from offset for nbytes with flags.
// It returns 0 on success and any error encountered.
func SyncFileRange(fd int, offset int64, nbytes int64, flags uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_sync_file_range), uintptr(fd), uintptr(offset), uintptr(nbytes), uintptr(flags), 0, 0)
	return r1, err
}

// Vmsplice splices user pages into a pipe.
// It splices uiov with nr_segs into fd with flags.
// It returns the number of bytes spliced and any error encountered.
func Vmsplice(fd int, uiov unsafe.Pointer, nrSegs uintptr, flags uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_vmsplice), uintptr(fd), uintptr(uiov), nrSegs, uintptr(flags), 0, 0)
	return r1, err
}

// MovePages moves pages between nodes.
// It moves nr_pages pages with nodes, status for pid with flags.
// It returns the number of pages not moved and any error encountered.
func MovePages(pid int, nrPages uintptr, pages unsafe.Pointer, nodes unsafe.Pointer, status unsafe.Pointer, flags int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_move_pages), uintptr(pid), nrPages, uintptr(pages), uintptr(nodes), uintptr(status), uintptr(flags))
	return r1, err
}

// Utimensat sets file timestamps with nanosecond precision.
// It sets utimes for filename relative to dfd with flags.
// It returns 0 on success and any error encountered.
func Utimensat(dfd int, filename []byte, utimes unsafe.Pointer, flags int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_utimensat), uintptr(dfd), uintptr(unsafe.Pointer(&filename[0])), uintptr(utimes), uintptr(flags), 0, 0)
	return r1, err
}

// EpollPwait waits for events on an epoll instance with signal mask.
// It waits on epfd for maxevents into events with timeout, sigmask, sigsetsize.
// It returns the number of events and any error encountered.
func EpollPwait(epfd int, events unsafe.Pointer, maxevents int, timeout int, sigmask unsafe.Pointer, sigsetsize uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_epoll_pwait), uintptr(epfd), uintptr(events), uintptr(maxevents), uintptr(timeout), uintptr(sigmask), sigsetsize)
	return r1, err
}

// Signalfd creates a file descriptor for accepting signals.
// It creates fd with user_mask of sizemask.
// It returns the descriptor and any error encountered.
func Signalfd(ufd int, userMask unsafe.Pointer, sizemask uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_signalfd), uintptr(ufd), uintptr(userMask), sizemask)
	return r1, err
}

// TimerfdCreate creates a timer file descriptor.
// It creates a timer with clockid, flags.
// It returns the descriptor and any error encountered.
func TimerfdCreate(clockid int, flags int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_timerfd_create), uintptr(clockid), uintptr(flags), 0)
	return r1, err
}

// Eventfd creates an event file descriptor.
// It creates an event with count.
// It returns the descriptor and any error encountered.
func Eventfd(count uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_eventfd), uintptr(count), 0, 0)
	return r1, err
}

// Fallocate allocates space for a file.
// It allocates mode space for fd from offset length.
// It returns 0 on success and any error encountered.
func Fallocate(fd int, mode int, offset int64, length int64) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_fallocate), uintptr(fd), uintptr(mode), uintptr(offset), uintptr(length), 0, 0)
	return r1, err
}

// TimerfdSettime sets the time of a timer file descriptor.
// It sets ufd with flags, utmr, otmr.
// It returns 0 on success and any error encountered.
func TimerfdSettime(ufd int, flags int, utmr unsafe.Pointer, otmr unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_timerfd_settime), uintptr(ufd), uintptr(flags), uintptr(utmr), uintptr(otmr), 0, 0)
	return r1, err
}

// TimerfdGettime gets the time of a timer file descriptor.
// It gets ufd into otmr.
// It returns 0 on success and any error encountered.
func TimerfdGettime(ufd int, otmr unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_timerfd_gettime), uintptr(ufd), uintptr(otmr), 0)
	return r1, err
}

// Accept4 accepts a connection on a socket.
// It accepts on fd into upeer_sockaddr upeer_addrlen with flags.
// It returns the descriptor and any error encountered.
func Accept4(fd int, upeerSockaddr unsafe.Pointer, upeerAddrlen unsafe.Pointer, flags int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_accept4), uintptr(fd), uintptr(upeerSockaddr), uintptr(upeerAddrlen), uintptr(flags), 0, 0)
	return r1, err
}

// Signalfd4 creates a file descriptor for accepting signals with flags.
// It creates ufd with user_mask of sizemask, flags.
// It returns the descriptor and any error encountered.
func Signalfd4(ufd int, userMask unsafe.Pointer, sizemask uintptr, flags int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_signalfd4), uintptr(ufd), uintptr(userMask), sizemask, uintptr(flags), 0, 0)
	return r1, err
}

// Eventfd2 creates an event file descriptor with flags.
// It creates an event with count, flags.
// It returns the descriptor and any error encountered.
func Eventfd2(count uint, flags int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_eventfd2), uintptr(count), uintptr(flags), 0)
	return r1, err
}

// EpollCreate1 creates an epoll instance with flags.
// It creates an epoll with flags.
// It returns the descriptor and any error encountered.
func EpollCreate1(flags int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_epoll_create1), uintptr(flags), 0, 0)
	return r1, err
}

// Dup3 duplicates a file descriptor with flags.
// It duplicates oldfd to newfd with flags.
// It returns the descriptor and any error encountered.
func Dup3(oldfd uint, newfd uint, flags int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_dup3), uintptr(oldfd), uintptr(newfd), uintptr(flags))
	return r1, err
}

// Pipe2 creates a pipe with flags.
// It creates a pipe into fildes with flags.
// It returns 0 on success and any error encountered.
func Pipe2(fildes unsafe.Pointer, flags int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_pipe2), uintptr(fildes), uintptr(flags), 0)
	return r1, err
}

// InotifyInit1 initializes an inotify instance with flags.
// It creates an inotify with flags.
// It returns the descriptor and any error encountered.
func InotifyInit1(flags int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_inotify_init1), uintptr(flags), 0, 0)
	return r1, err
}

// Preadv reads data into multiple buffers.
// It reads fd into vec vlen from pos_l pos_h.
// It returns the number of bytes read and any error encountered.
func Preadv(fd uintptr, vec unsafe.Pointer, vlen uintptr, posL uintptr, posH uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_preadv), fd, uintptr(vec), vlen, posL, posH, 0)
	return r1, err
}

// Pwritev writes data from multiple buffers.
// It writes fd from vec vlen to pos_l pos_h.
// It returns the number of bytes written and any error encountered.
func Pwritev(fd uintptr, vec unsafe.Pointer, vlen uintptr, posL uintptr, posH uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_pwritev), fd, uintptr(vec), vlen, posL, posH, 0)
	return r1, err
}

// RtTgsigqueueinfo queues a signal and data to a thread group.
// It sends sig with uinfo to tgid pid.
// It returns 0 on success and any error encountered.
func RtTgsigqueueinfo(tgid int, pid int, sig int, uinfo unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_rt_tgsigqueueinfo), uintptr(tgid), uintptr(pid), uintptr(sig), uintptr(uinfo), 0, 0)
	return r1, err
}

// PerfEventOpen sets up performance monitoring.
// It opens event with attr_uptr, pid, cpu, group_fd, flags.
// It returns the descriptor and any error encountered.
func PerfEventOpen(attrUptr unsafe.Pointer, pid int, cpu int, groupFd int, flags uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_perf_event_open), uintptr(attrUptr), uintptr(pid), uintptr(cpu), uintptr(groupFd), flags, 0)
	return r1, err
}

// Recvmmsg receives multiple messages.
// It receives vlen messages into mmsg with flags, timeout.
// It returns the number of messages and any error encountered.
func Recvmmsg(fd int, mmsg unsafe.Pointer, vlen uint, flags uint, timeout unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_recvmmsg), uintptr(fd), uintptr(mmsg), uintptr(vlen), uintptr(flags), uintptr(timeout), 0)
	return r1, err
}

// FanotifyInit initializes a fanotify group.
// It initializes with flags, event_f_flags.
// It returns the descriptor and any error encountered.
func FanotifyInit(flags uint, eventFFlags uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_fanotify_init), uintptr(flags), uintptr(eventFFlags), 0)
	return r1, err
}

// FanotifyMark adds or removes a fanotify mark.
// It marks fanotify_fd with flags, mask, dfd, pathname.
// It returns 0 on success and any error encountered.
func FanotifyMark(fanotifyFd int, flags uint, mask uint64, dfd int, pathname []byte) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_fanotify_mark), uintptr(fanotifyFd), uintptr(flags), uintptr(mask), uintptr(dfd), uintptr(unsafe.Pointer(&pathname[0])), 0)
	return r1, err
}

// Prlimit64 gets and sets resource limits.
// It gets/sets pid resource with new_rlim, old_rlim.
// It returns 0 on success and any error encountered.
func Prlimit64(pid int, resource uint, newRlim unsafe.Pointer, oldRlim unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_prlimit64), uintptr(pid), uintptr(resource), uintptr(newRlim), uintptr(oldRlim), 0, 0)
	return r1, err
}

// NameToHandleAt opens a file by handle.
// It opens name relative to dfd into handle, mnt_id with flag.
// It returns 0 on success and any error encountered.
func NameToHandleAt(dfd int, name []byte, handle unsafe.Pointer, mntId unsafe.Pointer, flag int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_name_to_handle_at), uintptr(dfd), uintptr(unsafe.Pointer(&name[0])), uintptr(handle), uintptr(mntId), uintptr(flag), 0)
	return r1, err
}

// OpenByHandleAt opens a file by handle.
// It opens handle with flags.
// It returns the descriptor and any error encountered.
func OpenByHandleAt(mountdirfd int, handle unsafe.Pointer, flags int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_open_by_handle_at), uintptr(mountdirfd), uintptr(handle), uintptr(flags))
	return r1, err
}

// ClockAdjtime adjusts the time of a clock.
// It adjusts which_clock with utx.
// It returns the state and any error encountered.
func ClockAdjtime(whichClock int, utx unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_clock_adjtime), uintptr(whichClock), uintptr(utx), 0)
	return r1, err
}

// Syncfs synchronizes a filesystem.
// It syncs the filesystem of fd.
// It returns 0 on success and any error encountered.
func Syncfs(fd int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_syncfs), uintptr(fd), 0, 0)
	return r1, err
}

// Sendmmsg sends multiple messages.
// It sends vlen messages from mmsg with flags.
// It returns the number of messages sent and any error encountered.
func Sendmmsg(fd int, mmsg unsafe.Pointer, vlen uint, flags uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_sendmmsg), uintptr(fd), uintptr(mmsg), uintptr(vlen), uintptr(flags), 0, 0)
	return r1, err
}

// Setns reassociates the calling thread with a namespace.
// It sets the namespace to fd with flags.
// It returns 0 on success and any error encountered.
func Setns(fd int, flags int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_setns), uintptr(fd), uintptr(flags), 0)
	return r1, err
}

// Getcpu gets the CPU and NUMA node.
// It gets cpup, nodep, unused.
// It returns 0 on success and any error encountered.
func Getcpu(cpup unsafe.Pointer, nodep unsafe.Pointer, unused unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_getcpu), uintptr(cpup), uintptr(nodep), uintptr(unused))
	return r1, err
}

// ProcessVmReadv reads memory from another process.
// It reads pid from lvec liovcnt into rvec riovcnt with flags.
// It returns the number of bytes read and any error encountered.
func ProcessVmReadv(pid int, lvec unsafe.Pointer, liovcnt uintptr, rvec unsafe.Pointer, riovcnt uintptr, flags uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_process_vm_readv), uintptr(pid), uintptr(lvec), liovcnt, uintptr(rvec), riovcnt, flags)
	return r1, err
}

// ProcessVmWritev writes memory to another process.
// It writes pid from lvec liovcnt to rvec riovcnt with flags.
// It returns the number of bytes written and any error encountered.
func ProcessVmWritev(pid int, lvec unsafe.Pointer, liovcnt uintptr, rvec unsafe.Pointer, riovcnt uintptr, flags uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_process_vm_writev), uintptr(pid), uintptr(lvec), liovcnt, uintptr(rvec), riovcnt, flags)
	return r1, err
}

// Kcmp compares two processes.
// It compares pid1 pid2 with type, idx1, idx2.
// It returns the result and any error encountered.
func Kcmp(pid1 int, pid2 int, typ int, idx1 uintptr, idx2 uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_kcmp), uintptr(pid1), uintptr(pid2), uintptr(typ), idx1, idx2, 0)
	return r1, err
}

// FinitModule loads a kernel module from a file.
// It loads fd with uargs, flags.
// It returns 0 on success and any error encountered.
func FinitModule(fd int, uargs []byte, flags int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_finit_module), uintptr(fd), uintptr(unsafe.Pointer(&uargs[0])), uintptr(flags))
	return r1, err
}

// SchedSetattr sets scheduling attributes.
// It sets pid with uattr, flags.
// It returns 0 on success and any error encountered.
func SchedSetattr(pid int, uattr unsafe.Pointer, flags uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_sched_setattr), uintptr(pid), uintptr(uattr), uintptr(flags))
	return r1, err
}

// SchedGetattr gets scheduling attributes.
// It gets pid into uattr with usize, flags.
// It returns 0 on success and any error encountered.
func SchedGetattr(pid int, uattr unsafe.Pointer, usize uint, flags uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_sched_getattr), uintptr(pid), uintptr(uattr), uintptr(usize), uintptr(flags), 0, 0)
	return r1, err
}

// Renameat2 renames a file with flags.
// It renames oldname to newname relative to olddfd/newdfd with flags.
// It returns 0 on success and any error encountered.
func Renameat2(olddfd int, oldname []byte, newdfd int, newname []byte, flags uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_renameat2), uintptr(olddfd), uintptr(unsafe.Pointer(&oldname[0])), uintptr(newdfd), uintptr(unsafe.Pointer(&newname[0])), uintptr(flags), 0)
	return r1, err
}

// Seccomp operates on the secure computing state.
// It performs op with flags, uargs.
// It returns 0 on success and any error encountered.
func Seccomp(op uint, flags uint, uargs unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_seccomp), uintptr(op), uintptr(flags), uintptr(uargs))
	return r1, err
}

// Getrandom gets random bytes.
// It gets len bytes into ubuf with flags.
// It returns the number of bytes and any error encountered.
func Getrandom(ubuf unsafe.Pointer, length uintptr, flags uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_getrandom), uintptr(ubuf), length, uintptr(flags))
	return r1, err
}

// MemfdCreate creates an anonymous file.
// It creates uname with flags.
// It returns the descriptor and any error encountered.
func MemfdCreate(uname []byte, flags uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_memfd_create), uintptr(unsafe.Pointer(&uname[0])), uintptr(flags), 0)
	return r1, err
}

// KexecFileLoad loads a new kernel from a file.
// It loads kernel_fd, initrd_fd with cmdline_len, cmdline_ptr, flags.
// It returns 0 on success and any error encountered.
func KexecFileLoad(kernelFd int, initrdFd int, cmdlineLen uintptr, cmdlinePtr []byte, flags uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_kexec_file_load), uintptr(kernelFd), uintptr(initrdFd), cmdlineLen, uintptr(unsafe.Pointer(&cmdlinePtr[0])), flags, 0)
	return r1, err
}

// Bpf performs a BPF operation.
// It performs cmd with uattr, size.
// It returns the result and any error encountered.
func Bpf(cmd int, uattr unsafe.Pointer, size uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_bpf), uintptr(cmd), uintptr(uattr), uintptr(size))
	return r1, err
}

// Execveat executes a program relative to a directory.
// It executes fd filename with argv, envp, flags.
// It returns 0 on success and any error encountered.
func Execveat(fd int, filename []byte, argv unsafe.Pointer, envp unsafe.Pointer, flags int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_execveat), uintptr(fd), uintptr(unsafe.Pointer(&filename[0])), uintptr(argv), uintptr(envp), uintptr(flags), 0)
	return r1, err
}

// Userfaultfd creates a userfaultfd object.
// It creates with flags.
// It returns the descriptor and any error encountered.
func Userfaultfd(flags int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_userfaultfd), uintptr(flags), 0, 0)
	return r1, err
}

// Membarrier performs a memory barrier.
// It performs cmd with flags, cpu_id.
// It returns 0 on success and any error encountered.
func Membarrier(cmd int, flags uint, cpuId int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_membarrier), uintptr(cmd), uintptr(flags), uintptr(cpuId))
	return r1, err
}

// Mlock2 locks pages in memory with flags.
// It locks len bytes starting at start with flags.
// It returns 0 on success and any error encountered.
func Mlock2(start uintptr, length uintptr, flags int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_mlock2), start, length, uintptr(flags))
	return r1, err
}

// CopyFileRange copies data between files.
// It copies len bytes from fd_in off_in to fd_out off_out with flags.
// It returns the number of bytes copied and any error encountered.
func CopyFileRange(fdIn int, offIn unsafe.Pointer, fdOut int, offOut unsafe.Pointer, length uintptr, flags uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_copy_file_range), uintptr(fdIn), uintptr(offIn), uintptr(fdOut), uintptr(offOut), length, uintptr(flags))
	return r1, err
}

// Preadv2 reads data into multiple buffers with flags.
// It reads fd into vec vlen from pos_l pos_h with flags.
// It returns the number of bytes read and any error encountered.
func Preadv2(fd uintptr, vec unsafe.Pointer, vlen uintptr, posL uintptr, posH uintptr, flags int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_preadv2), fd, uintptr(vec), vlen, posL, posH, uintptr(flags))
	return r1, err
}

// Pwritev2 writes data from multiple buffers with flags.
// It writes fd from vec vlen to pos_l pos_h with flags.
// It returns the number of bytes written and any error encountered.
func Pwritev2(fd uintptr, vec unsafe.Pointer, vlen uintptr, posL uintptr, posH uintptr, flags int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_pwritev2), fd, uintptr(vec), vlen, posL, posH, uintptr(flags))
	return r1, err
}

// PkeyMprotect sets memory protection with a protection key.
// It protects start len with prot, pkey.
// It returns 0 on success and any error encountered.
func PkeyMprotect(start uintptr, length uintptr, prot uintptr, pkey int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_pkey_mprotect), start, length, prot, uintptr(pkey), 0, 0)
	return r1, err
}

// PkeyAlloc allocates a protection key.
// It allocates with flags, init_val.
// It returns the key and any error encountered.
func PkeyAlloc(flags uintptr, initVal uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_pkey_alloc), flags, initVal, 0)
	return r1, err
}

// PkeyFree frees a protection key.
// It frees pkey.
// It returns 0 on success and any error encountered.
func PkeyFree(pkey int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_pkey_free), uintptr(pkey), 0, 0)
	return r1, err
}

// Statx gets extended file status.
// It gets statx of filename relative to dfd with flags, mask into buffer.
// It returns 0 on success and any error encountered.
func Statx(dfd int, filename []byte, flags uint, mask uint, buffer unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_statx), uintptr(dfd), uintptr(unsafe.Pointer(&filename[0])), uintptr(flags), uintptr(mask), uintptr(buffer), 0)
	return r1, err
}

// IoPgetevents reads asynchronous I/O events with signal mask.
// It reads ctx_id min_nr nr into events, timeout, usig.
// It returns the number of events and any error encountered.
func IoPgetevents(ctxId uintptr, minNr int64, nr int64, events unsafe.Pointer, timeout unsafe.Pointer, usig unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_io_pgetevents), ctxId, uintptr(minNr), uintptr(nr), uintptr(events), uintptr(timeout), uintptr(usig))
	return r1, err
}

// Rseq registers a restartable sequence.
// It registers rseq with rseq_len, flags, sig.
// It returns 0 on success and any error encountered.
func Rseq(rseq unsafe.Pointer, rseqLen uint32, flags int, sig uint32) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_rseq), uintptr(rseq), uintptr(rseqLen), uintptr(flags), uintptr(sig), 0, 0)
	return r1, err
}

// Uretprobe is not implemented.
// It is a placeholder for uretprobe.
// It returns -ENOSYS.
func Uretprobe() (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_uretprobe), 0, 0, 0)
	return r1, err
}

// PidfdSendSignal sends a signal to a process via pidfd.
// It sends sig with info, flags to pidfd.
// It returns 0 on success and any error encountered.
func PidfdSendSignal(pidfd int, sig int, info unsafe.Pointer, flags uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_pidfd_send_signal), uintptr(pidfd), uintptr(sig), uintptr(info), uintptr(flags), 0, 0)
	return r1, err
}

// IoUringSetup sets up an io_uring instance.
// It sets up entries into params.
// It returns the descriptor and any error encountered.
func IoUringSetup(entries uint32, params unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_io_uring_setup), uintptr(entries), uintptr(params), 0)
	return r1, err
}

// IoUringEnter initiates and completes I/O operations.
// It enters fd with to_submit, min_complete, flags, argp, argsz.
// It returns the result and any error encountered.
func IoUringEnter(fd uint, toSubmit uint32, minComplete uint32, flags uint32, argp unsafe.Pointer, argsz uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_io_uring_enter), uintptr(fd), uintptr(toSubmit), uintptr(minComplete), uintptr(flags), uintptr(argp), argsz)
	return r1, err
}

// IoUringRegister registers files or user buffers.
// It registers fd with opcode, arg, nr_args.
// It returns 0 on success and any error encountered.
func IoUringRegister(fd uint, opcode uint, arg unsafe.Pointer, nrArgs uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_io_uring_register), uintptr(fd), uintptr(opcode), uintptr(arg), uintptr(nrArgs), 0, 0)
	return r1, err
}

// OpenTree opens the root of a mount tree.
// It opens filename relative to dfd with flags.
// It returns the descriptor and any error encountered.
func OpenTree(dfd int, filename []byte, flags uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_open_tree), uintptr(dfd), uintptr(unsafe.Pointer(&filename[0])), uintptr(flags))
	return r1, err
}

// MoveMount moves a mount.
// It moves from_pathname to to_pathname relative to from_dfd to_dfd with flags.
// It returns 0 on success and any error encountered.
func MoveMount(fromDfd int, fromPathname []byte, toDfd int, toPathname []byte, flags uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_move_mount), uintptr(fromDfd), uintptr(unsafe.Pointer(&fromPathname[0])), uintptr(toDfd), uintptr(unsafe.Pointer(&toPathname[0])), uintptr(flags), 0)
	return r1, err
}

// Fsopen opens a filesystem.
// It opens _fs_name with flags.
// It returns the descriptor and any error encountered.
func Fsopen(_fsName []byte, flags uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_fsopen), uintptr(unsafe.Pointer(&_fsName[0])), uintptr(flags), 0)
	return r1, err
}

// Fsconfig configures a filesystem.
// It configures fd with cmd, _key, _value, aux.
// It returns 0 on success and any error encountered.
func Fsconfig(fd int, cmd uint, _key []byte, _value unsafe.Pointer, aux int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_fsconfig), uintptr(fd), uintptr(cmd), uintptr(unsafe.Pointer(&_key[0])), uintptr(_value), uintptr(aux), 0)
	return r1, err
}

// Fsmount mounts a filesystem.
// It mounts fs_fd with flags, attr_flags.
// It returns the descriptor and any error encountered.
func Fsmount(fsFd int, flags uint, attrFlags uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_fsmount), uintptr(fsFd), uintptr(flags), uintptr(attrFlags))
	return r1, err
}

// Fspick picks a filesystem.
// It picks path relative to dfd with flags.
// It returns the descriptor and any error encountered.
func Fspick(dfd int, path []byte, flags uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_fspick), uintptr(dfd), uintptr(unsafe.Pointer(&path[0])), uintptr(flags))
	return r1, err
}

// PidfdOpen opens a process file descriptor.
// It opens pid with flags.
// It returns the descriptor and any error encountered.
func PidfdOpen(pid int, flags uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_pidfd_open), uintptr(pid), uintptr(flags), 0)
	return r1, err
}

// Clone3 creates a child process.
// It creates with uargs, size.
// It returns the pid and any error encountered.
func Clone3(uargs unsafe.Pointer, size uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_clone3), uintptr(uargs), size, 0)
	return r1, err
}

// CloseRange closes a range of file descriptors.
// It closes from fd to max_fd with flags.
// It returns 0 on success and any error encountered.
func CloseRange(fd uint, maxFd uint, flags uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_close_range), uintptr(fd), uintptr(maxFd), uintptr(flags))
	return r1, err
}

// Openat2 opens a file with extended options.
// It opens filename relative to dfd with how, usize.
// It returns the descriptor and any error encountered.
func Openat2(dfd int, filename []byte, how unsafe.Pointer, usize uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_openat2), uintptr(dfd), uintptr(unsafe.Pointer(&filename[0])), uintptr(how), usize, 0, 0)
	return r1, err
}

// PidfdGetfd gets a file descriptor from a process.
// It gets fd from pidfd with flags.
// It returns the descriptor and any error encountered.
func PidfdGetfd(pidfd int, fd int, flags uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_pidfd_getfd), uintptr(pidfd), uintptr(fd), uintptr(flags))
	return r1, err
}

// Faccessat2 checks access to a file with flags.
// It checks mode access to filename relative to dfd with flags.
// It returns 0 on success and any error encountered.
func Faccessat2(dfd int, filename []byte, mode int, flags int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_faccessat2), uintptr(dfd), uintptr(unsafe.Pointer(&filename[0])), uintptr(mode), uintptr(flags))
	return r1, err
}

// ProcessMadvise gives advice about memory usage.
// It advises pidfd with vec, vlen, behavior, flags.
// It returns the result and any error encountered.
func ProcessMadvise(pidfd int, vec unsafe.Pointer, vlen uintptr, behavior int, flags uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_process_madvise), uintptr(pidfd), uintptr(vec), vlen, uintptr(behavior), uintptr(flags), 0)
	return r1, err
}

// EpollPwait2 waits for events on an epoll instance with timeout.
// It waits on epfd for maxevents into events with timeout, sigmask, sigsetsize.
// It returns the number of events and any error encountered.
func EpollPwait2(epfd int, events unsafe.Pointer, maxevents int, timeout unsafe.Pointer, sigmask unsafe.Pointer, sigsetsize uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_epoll_pwait2), uintptr(epfd), uintptr(events), uintptr(maxevents), uintptr(timeout), uintptr(sigmask), sigsetsize)
	return r1, err
}

// MountSetattr sets mount attributes.
// It sets path relative to dfd with flags, uattr, usize.
// It returns 0 on success and any error encountered.
func MountSetattr(dfd int, path []byte, flags uint, uattr unsafe.Pointer, usize uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_mount_setattr), uintptr(dfd), uintptr(unsafe.Pointer(&path[0])), uintptr(flags), uintptr(uattr), usize, 0)
	return r1, err
}

// QuotactlFd manipulates disk quotas by file descriptor.
// It performs cmd on fd for id with addr.
// It returns 0 on success and any error encountered.
func QuotactlFd(fd uint, cmd uint, id int, addr unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_quotactl_fd), uintptr(fd), uintptr(cmd), uintptr(id), uintptr(addr), 0, 0)
	return r1, err
}

// LandlockCreateRuleset creates a Landlock ruleset.
// It creates with attr, size, flags.
// It returns the descriptor and any error encountered.
func LandlockCreateRuleset(attr unsafe.Pointer, size uintptr, flags uint32) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_landlock_create_ruleset), uintptr(attr), size, uintptr(flags))
	return r1, err
}

// LandlockAddRule adds a rule to a Landlock ruleset.
// It adds rule_type rule_attr to ruleset_fd with flags.
// It returns 0 on success and any error encountered.
func LandlockAddRule(rulesetFd int, ruleType int, ruleAttr unsafe.Pointer, flags uint32) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_landlock_add_rule), uintptr(rulesetFd), uintptr(ruleType), uintptr(ruleAttr), uintptr(flags), 0, 0)
	return r1, err
}

// LandlockRestrictSelf restricts the current process.
// It restricts with ruleset_fd, flags.
// It returns 0 on success and any error encountered.
func LandlockRestrictSelf(rulesetFd int, flags uint32) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_landlock_restrict_self), uintptr(rulesetFd), uintptr(flags), 0)
	return r1, err
}

// MemfdSecret creates a secret memory file.
// It creates with flags.
// It returns the descriptor and any error encountered.
func MemfdSecret(flags uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_memfd_secret), uintptr(flags), 0, 0)
	return r1, err
}

// ProcessMrelease releases process memory.
// It releases pidfd with flags.
// It returns 0 on success and any error encountered.
func ProcessMrelease(pidfd int, flags uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_process_mrelease), uintptr(pidfd), uintptr(flags), 0)
	return r1, err
}

// FutexWaitv waits on multiple futexes.
// It waits on waiters nr_futexes with flags, timeout, clockid.
// It returns 0 on success and any error encountered.
func FutexWaitv(waiters unsafe.Pointer, nrFutexes uint, flags uint, timeout unsafe.Pointer, clockid int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_futex_waitv), uintptr(waiters), uintptr(nrFutexes), uintptr(flags), uintptr(timeout), uintptr(clockid), 0)
	return r1, err
}

// SetMempolicyHomeNode sets the home node for memory policy.
// It sets start length to home_node with flags.
// It returns 0 on success and any error encountered.
func SetMempolicyHomeNode(start uintptr, length uintptr, homeNode uintptr, flags uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_set_mempolicy_home_node), start, length, homeNode, flags, 0, 0)
	return r1, err
}

// Cachestat gets cache statistics.
// It gets fd cstat_range into cstat with flags.
// It returns 0 on success and any error encountered.
func Cachestat(fd uint, cstatRange unsafe.Pointer, cstat unsafe.Pointer, flags uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_cachestat), uintptr(fd), uintptr(cstatRange), uintptr(cstat), uintptr(flags), 0, 0)
	return r1, err
}

// Fchmodat2 changes permissions of a file with flags.
// It changes mode of filename relative to dfd with flags.
// It returns 0 on success and any error encountered.
func Fchmodat2(dfd int, filename []byte, mode uint, flags uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_fchmodat2), uintptr(dfd), uintptr(unsafe.Pointer(&filename[0])), uintptr(mode), uintptr(flags), 0, 0)
	return r1, err
}

// MapShadowStack maps a shadow stack.
// It maps addr size with flags.
// It returns the address and any error encountered.
func MapShadowStack(addr uintptr, size uintptr, flags uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_map_shadow_stack), addr, size, uintptr(flags))
	return r1, err
}

// FutexWake wakes waiters on a futex.
// It wakes uaddr with mask, nr, flags.
// It returns the number woken and any error encountered.
func FutexWake(uaddr unsafe.Pointer, mask uintptr, nr int, flags uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_futex_wake), uintptr(uaddr), mask, uintptr(nr), uintptr(flags), 0, 0)
	return r1, err
}

// FutexWait waits on a futex.
// It waits on uaddr val with mask, flags, timeout, clockid.
// It returns 0 on success and any error encountered.
func FutexWait(uaddr unsafe.Pointer, val uintptr, mask uintptr, flags uint, timeout unsafe.Pointer, clockid int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_futex_wait), uintptr(uaddr), val, mask, uintptr(flags), uintptr(timeout), uintptr(clockid))
	return r1, err
}

// FutexRequeue requeues waiters on a futex.
// It requeues waiters with nr_wake, nr_requeue.
// It returns the number requeued and any error encountered.
func FutexRequeue(waiters unsafe.Pointer, flags int, nrWake int, nrRequeue int) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_futex_requeue), uintptr(waiters), uintptr(flags), uintptr(nrWake), uintptr(nrRequeue), 0, 0)
	return r1, err
}

// Statmount gets mount information.
// It gets req into buf bufsize with flags.
// It returns 0 on success and any error encountered.
func Statmount(req unsafe.Pointer, buf unsafe.Pointer, bufsize uintptr, flags uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_statmount), uintptr(req), uintptr(buf), bufsize, uintptr(flags), 0, 0)
	return r1, err
}

// Listmount lists mounts.
// It lists req into mnt_ids nr_mnt_ids with flags.
// It returns the number listed and any error encountered.
func Listmount(req unsafe.Pointer, mntIds unsafe.Pointer, nrMntIds uintptr, flags uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_listmount), uintptr(req), uintptr(mntIds), nrMntIds, uintptr(flags), 0, 0)
	return r1, err
}

// LsmGetSelfAttr gets LSM attributes.
// It gets attr into ctx size with flags.
// It returns 0 on success and any error encountered.
func LsmGetSelfAttr(attr uint, ctx unsafe.Pointer, size unsafe.Pointer, flags uint32) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_lsm_get_self_attr), uintptr(attr), uintptr(ctx), uintptr(size), uintptr(flags), 0, 0)
	return r1, err
}

// LsmSetSelfAttr sets LSM attributes.
// It sets attr ctx size with flags.
// It returns 0 on success and any error encountered.
func LsmSetSelfAttr(attr uint, ctx unsafe.Pointer, size uint32, flags uint32) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_lsm_set_self_attr), uintptr(attr), uintptr(ctx), uintptr(size), uintptr(flags), 0, 0)
	return r1, err
}

// LsmListModules lists LSM modules.
// It lists into ids size with flags.
// It returns 0 on success and any error encountered.
func LsmListModules(ids unsafe.Pointer, size unsafe.Pointer, flags uint32) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_lsm_list_modules), uintptr(ids), uintptr(size), uintptr(flags))
	return r1, err
}

// Mseal seals memory.
// It seals start len with flags.
// It returns 0 on success and any error encountered.
func Mseal(start uintptr, length uintptr, flags uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_mseal), start, length, flags)
	return r1, err
}

// Setxattrat sets an extended attribute with flags.
// It sets pathname at_flags name uargs usize relative to dfd.
// It returns 0 on success and any error encountered.
func Setxattrat(dfd int, pathname []byte, atFlags uint, name []byte, uargs unsafe.Pointer, usize uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_setxattrat), uintptr(dfd), uintptr(unsafe.Pointer(&pathname[0])), uintptr(atFlags), uintptr(unsafe.Pointer(&name[0])), uintptr(uargs), usize)
	return r1, err
}

// Getxattrat gets an extended attribute with flags.
// It gets pathname at_flags name uargs usize relative to dfd.
// It returns the size and any error encountered.
func Getxattrat(dfd int, pathname []byte, atFlags uint, name []byte, uargs unsafe.Pointer, usize uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_getxattrat), uintptr(dfd), uintptr(unsafe.Pointer(&pathname[0])), uintptr(atFlags), uintptr(unsafe.Pointer(&name[0])), uintptr(uargs), usize)
	return r1, err
}

// Listxattrat lists extended attributes with flags.
// It lists pathname at_flags into list size relative to dfd.
// It returns the size and any error encountered.
func Listxattrat(dfd int, pathname []byte, atFlags uint, list unsafe.Pointer, size uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_listxattrat), uintptr(dfd), uintptr(unsafe.Pointer(&pathname[0])), uintptr(atFlags), uintptr(list), size, 0)
	return r1, err
}

// Removexattrat removes an extended attribute with flags.
// It removes pathname at_flags name relative to dfd.
// It returns 0 on success and any error encountered.
func Removexattrat(dfd int, pathname []byte, atFlags uint, name []byte) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_removexattrat), uintptr(dfd), uintptr(unsafe.Pointer(&pathname[0])), uintptr(atFlags), uintptr(unsafe.Pointer(&name[0])), 0, 0)
	return r1, err
}

// OpenTreeAttr opens the root of a mount tree with attributes.
// It opens filename relative to dfd with flags, uattr, usize.
// It returns the descriptor and any error encountered.
func OpenTreeAttr(dfd int, filename []byte, flags uint, uattr unsafe.Pointer, usize uintptr) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_open_tree_attr), uintptr(dfd), uintptr(unsafe.Pointer(&filename[0])), uintptr(flags), uintptr(uattr), usize, 0)
	return r1, err
}

// FileGetattr gets file attributes.
// It gets filename relative to dfd into ufattr usize with at_flags.
// It returns 0 on success and any error encountered.
func FileGetattr(dfd int, filename []byte, ufattr unsafe.Pointer, usize uintptr, atFlags uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_file_getattr), uintptr(dfd), uintptr(unsafe.Pointer(&filename[0])), uintptr(ufattr), usize, uintptr(atFlags), 0)
	return r1, err
}

// FileSetattr sets file attributes.
// It sets filename relative to dfd with ufattr usize at_flags.
// It returns 0 on success and any error encountered.
func FileSetattr(dfd int, filename []byte, ufattr unsafe.Pointer, usize uintptr, atFlags uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_file_setattr), uintptr(dfd), uintptr(unsafe.Pointer(&filename[0])), uintptr(ufattr), usize, uintptr(atFlags), 0)
	return r1, err
}
