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
	__x64_sys_read = 0
	__x64_sys_write = 1
	__x64_sys_open = 2
	__x64_sys_close = 3
	__x64_sys_newstat = 4
	__x64_sys_newfstat = 5
	__x64_sys_newlstat = 6
	__x64_sys_poll = 7
	__x64_sys_lseek = 8
	__x64_sys_mmap = 9
	__x64_sys_mprotect = 10
	__x64_sys_munmap = 11
	__x64_sys_brk = 12
	__x64_sys_rt_sigaction = 13
	__x64_sys_rt_sigprocmask = 14
	__x64_sys_rt_sigreturn = 15
	__x64_sys_ioctl = 16
	__x64_sys_pread64 = 17
	__x64_sys_pwrite64 = 18
	__x64_sys_readv = 19
	__x64_sys_writev = 20
	__x64_sys_access = 21
	__x64_sys_pipe = 22
	__x64_sys_select = 23
	__x64_sys_sched_yield = 24
	__x64_sys_mremap = 25
	__x64_sys_msync = 26
	__x64_sys_mincore = 27
	__x64_sys_madvise = 28
	__x64_sys_shmget = 29
	__x64_sys_shmat = 30
	__x64_sys_shmctl = 31
	__x64_sys_dup = 32
	__x64_sys_dup2 = 33
	__x64_sys_pause = 34
	__x64_sys_nanosleep = 35
	__x64_sys_getitimer = 36
	__x64_sys_alarm = 37
	__x64_sys_setitimer = 38
	__x64_sys_getpid = 39
	__x64_sys_sendfile64
	__x64_sys_socket = 41
	__x64_sys_connect = 42
	__x64_sys_accept = 43
	__x64_sys_sendto = 44
	__x64_sys_recvfrom = 45
	__x64_sys_sendmsg = 46
	__x64_sys_recvmsg = 47
	__x64_sys_shutdown = 48
	__x64_sys_bind = 49
	__x64_sys_listen = 50
	__x64_sys_getsockname = 51
	__x64_sys_getpeername = 52
	__x64_sys_socketpair = 53
	__x64_sys_setsockopt = 54
	__x64_sys_getsockopt = 55
	__x64_sys_clone = 56
	__x64_sys_fork = 57
	__x64_sys_vfork = 58
	__x64_sys_execve = 59
	__x64_sys_exit = 60
	__x64_sys_wait4 = 61
	__x64_sys_kill = 62
	__x64_sys_newuname
	__x64_sys_semget = 64
	__x64_sys_semop = 65
	__x64_sys_semctl = 66
	__x64_sys_shmdt = 67
	__x64_sys_msgget = 68
	__x64_sys_msgsnd = 69
	__x64_sys_msgrcv = 70
	__x64_sys_msgctl = 71
	__x64_sys_fcntl = 72
	__x64_sys_flock = 73
	__x64_sys_fsync = 74
	__x64_sys_fdatasync = 75
	__x64_sys_truncate = 76
	__x64_sys_ftruncate = 77
	__x64_sys_getdents = 78
	__x64_sys_getcwd = 79
	__x64_sys_chdir = 80
	__x64_sys_fchdir = 81
	__x64_sys_rename = 82
	__x64_sys_mkdir = 83
	__x64_sys_rmdir = 84
	__x64_sys_creat = 85
	__x64_sys_link = 86
	__x64_sys_unlink = 87
	__x64_sys_symlink = 88
	__x64_sys_readlink = 89
	__x64_sys_chmod = 90
	__x64_sys_fchmod = 91
	__x64_sys_chown = 92
	__x64_sys_fchown = 93
	__x64_sys_lchown = 94
	__x64_sys_umask = 95
	__x64_sys_gettimeofday = 96
	__x64_sys_getrlimit = 97
	__x64_sys_getrusage = 98
	__x64_sys_sysinfo = 99
	__x64_sys_times = 100
	__x64_sys_ptrace = 101
	__x64_sys_getuid = 102
	__x64_sys_syslog = 103
	__x64_sys_getgid = 104
	__x64_sys_setuid = 105
	__x64_sys_setgid = 106
	__x64_sys_geteuid = 107
	__x64_sys_getegid = 108
	__x64_sys_setpgid = 109
	__x64_sys_getppid = 110
	__x64_sys_getpgrp = 111
	__x64_sys_setsid = 112
	__x64_sys_setreuid = 113
	__x64_sys_setregid = 114
	__x64_sys_getgroups = 115
	__x64_sys_setgroups = 116
	__x64_sys_setresuid = 117
	__x64_sys_getresuid = 118
	__x64_sys_setresgid = 119
	__x64_sys_getresgid = 120
	__x64_sys_getpgid = 121
	__x64_sys_setfsuid = 122
	__x64_sys_setfsgid = 123
	__x64_sys_getsid = 124
	__x64_sys_capget = 125
	__x64_sys_capset = 126
	__x64_sys_rt_sigpending = 127
	__x64_sys_rt_sigtimedwait = 128
	__x64_sys_rt_sigqueueinfo = 129
	__x64_sys_rt_sigsuspend = 130
	__x64_sys_sigaltstack = 131
	__x64_sys_utime = 132
	__x64_sys_mknod = 133
	__x64_sys_personality = 135
	__x64_sys_ustat = 136
	__x64_sys_statfs = 137
	__x64_sys_fstatfs = 138
	__x64_sys_sysfs = 139
	__x64_sys_getpriority = 140
	__x64_sys_setpriority = 141
	__x64_sys_sched_setparam = 142
	__x64_sys_sched_getparam = 143
	__x64_sys_sched_setscheduler = 144
	__x64_sys_sched_getscheduler = 145
	__x64_sys_sched_get_priority_max = 146
	__x64_sys_sched_get_priority_min = 147
	__x64_sys_sched_rr_get_interval = 148
	__x64_sys_mlock = 149
	__x64_sys_munlock = 150
	__x64_sys_mlockall = 151
	__x64_sys_munlockall = 152
	__x64_sys_vhangup = 153
	__x64_sys_modify_ldt = 154
	__x64_sys_pivot_root = 155
	__x64_sys_prctl = 156
	__x64_sys_arch_prctl = 157
	__x64_sys_adjtimex = 158
	__x64_sys_setrlimit = 159
	__x64_sys_chroot = 160
	__x64_sys_sync = 161
	__x64_sys_acct = 162
	__x64_sys_settimeofday = 163
	__x64_sys_mount = 164
	__x64_sys_umount = 165
	__x64_sys_swapon = 166
	__x64_sys_swapoff = 167
	__x64_sys_reboot = 168
	__x64_sys_sethostname = 169
	__x64_sys_setdomainname = 170
	__x64_sys_iopl = 171
	__x64_sys_ioperm = 172
	__x64_sys_init_module = 174
	__x64_sys_delete_module = 175
	__x64_sys_quotactl = 178
	__x64_sys_gettid = 185
	__x64_sys_readahead = 186
	__x64_sys_setxattr = 187
	__x64_sys_lsetxattr = 188
	__x64_sys_fsetxattr = 189
	__x64_sys_getxattr = 190
	__x64_sys_lgetxattr = 191
	__x64_sys_fgetxattr = 192
	__x64_sys_listxattr = 193
	__x64_sys_llistxattr = 194
	__x64_sys_flistxattr = 195
	__x64_sys_removexattr = 196
	__x64_sys_lremovexattr = 197
	__x64_sys_fremovexattr = 198
	__x64_sys_tkill = 199
	__x64_sys_time = 200
	__x64_sys_futex = 201
	__x64_sys_sched_setaffinity = 202
	__x64_sys_sched_getaffinity = 203
	__x64_sys_io_setup = 205
	__x64_sys_io_destroy = 206
	__x64_sys_io_getevents = 207
	__x64_sys_io_submit = 208
	__x64_sys_io_cancel = 209
	__x64_sys_epoll_create = 212
	__x64_sys_remap_file_pages = 215
	__x64_sys_getdents64 = 216
	__x64_sys_set_tid_address = 217
	__x64_sys_restart_syscall = 218
	__x64_sys_semtimedop = 219
	__x64_sys_fadvise64 = 220
	__x64_sys_timer_create = 221
	__x64_sys_timer_settime = 222
	__x64_sys_timer_gettime = 223
	__x64_sys_timer_getoverrun = 224
	__x64_sys_timer_delete = 225
	__x64_sys_clock_settime = 226
	__x64_sys_clock_gettime = 227
	__x64_sys_clock_getres = 228
	__x64_sys_clock_nanosleep = 229
	__x64_sys_exit_group = 230
	__x64_sys_epoll_wait = 231
	__x64_sys_epoll_ctl = 232
	__x64_sys_tgkill = 233
	__x64_sys_utimes = 234
	__x64_sys_mbind = 236
	__x64_sys_set_mempolicy = 237
	__x64_sys_get_mempolicy = 238
	__x64_sys_mq_open = 239
	__x64_sys_mq_unlink = 240
	__x64_sys_mq_timedsend = 241
	__x64_sys_mq_timedreceive = 242
	__x64_sys_mq_notify = 243
	__x64_sys_mq_getsetattr = 244
	__x64_sys_kexec_load = 245
	__x64_sys_waitid = 246
	__x64_sys_add_key = 247
	__x64_sys_request_key = 248
	__x64_sys_keyctl = 249
	__x64_sys_ioprio_set = 250
	__x64_sys_ioprio_get = 251
	__x64_sys_inotify_init = 252
	__x64_sys_inotify_add_watch = 253
	__x64_sys_inotify_rm_watch = 254
	__x64_sys_migrate_pages = 255
	__x64_sys_openat = 256
	__x64_sys_mkdirat = 257
	__x64_sys_mknodat = 258
	__x64_sys_fchownat = 259
	__x64_sys_futimesat = 260
	__x64_sys_newfstatat = 261
	__x64_sys_unlinkat = 262
	__x64_sys_renameat = 263
	__x64_sys_linkat = 264
	__x64_sys_symlinkat = 265
	__x64_sys_readlinkat = 266
	__x64_sys_fchmodat = 267
	__x64_sys_faccessat = 268
	__x64_sys_pselect6 = 269
	__x64_sys_ppoll = 270
	__x64_sys_unshare = 271
	__x64_sys_set_robust_list = 272
	__x64_sys_get_robust_list = 273
	__x64_sys_splice = 274
	__x64_sys_tee = 275
	__x64_sys_sync_file_range = 276
	__x64_sys_vmsplice = 277
	__x64_sys_move_pages = 278
	__x64_sys_utimensat = 279
	__x64_sys_epoll_pwait = 280
	__x64_sys_signalfd = 281
	__x64_sys_timerfd_create = 282
	__x64_sys_eventfd = 283
	__x64_sys_fallocate = 284
	__x64_sys_timerfd_settime = 285
	__x64_sys_timerfd_gettime = 286
	__x64_sys_accept4 = 287
	__x64_sys_signalfd4 = 288
	__x64_sys_eventfd2 = 289
	__x64_sys_epoll_create1 = 290
	__x64_sys_dup3 = 291
	__x64_sys_pipe2 = 292
	__x64_sys_inotify_init1 = 293
	__x64_sys_preadv = 294
	__x64_sys_pwritev = 295
	__x64_sys_rt_tgsigqueueinfo = 296
	__x64_sys_perf_event_open = 297
	__x64_sys_recvmmsg = 298
	__x64_sys_fanotify_init = 299
	__x64_sys_fanotify_mark = 300
	__x64_sys_prlimit64 = 301
	__x64_sys_name_to_handle_at = 302
	__x64_sys_open_by_handle_at = 303
	__x64_sys_clock_adjtime = 304
	__x64_sys_syncfs = 305
	__x64_sys_sendmmsg = 306
	__x64_sys_setns = 307
	__x64_sys_getcpu = 308
	__x64_sys_process_vm_readv = 309
	__x64_sys_process_vm_writev = 310
	__x64_sys_kcmp = 311
	__x64_sys_finit_module = 312
	__x64_sys_sched_setattr = 313
	__x64_sys_sched_getattr = 314
	__x64_sys_renameat2 = 315
	__x64_sys_seccomp = 316
	__x64_sys_getrandom = 317
	__x64_sys_memfd_create = 318
	__x64_sys_kexec_file_load = 319
	__x64_sys_bpf = 320
	__x64_sys_execveat = 321
	__x64_sys_userfaultfd = 322
	__x64_sys_membarrier = 323
	__x64_sys_mlock2 = 324
	__x64_sys_copy_file_range = 325
	__x64_sys_preadv2 = 326
	__x64_sys_pwritev2 = 327
	__x64_sys_pkey_mprotect = 328
	__x64_sys_pkey_alloc = 329
	__x64_sys_pkey_free = 330
	__x64_sys_statx = 331
	__x64_sys_io_pgetevents = 332
	__x64_sys_rseq = 333
	__x64_sys_uretprobe
	__x64_sys_pidfd_send_signal = 334
	__x64_sys_io_uring_setup = 335
	__x64_sys_io_uring_enter = 336
	__x64_sys_io_uring_register = 337
	__x64_sys_open_tree = 338
	__x64_sys_move_mount = 339
	__x64_sys_fsopen = 340
	__x64_sys_fsconfig = 341
	__x64_sys_fsmount = 342
	__x64_sys_fspick = 343
	__x64_sys_pidfd_open = 344
	__x64_sys_clone3 = 345
	__x64_sys_close_range = 346
	__x64_sys_openat2 = 347
	__x64_sys_pidfd_getfd = 348
	__x64_sys_faccessat2 = 349
	__x64_sys_process_madvise = 350
	__x64_sys_epoll_pwait2 = 351
	__x64_sys_mount_setattr = 352
	__x64_sys_quotactl_fd = 353
	__x64_sys_landlock_create_ruleset = 354
	__x64_sys_landlock_add_rule = 355
	__x64_sys_landlock_restrict_self = 356
	__x64_sys_memfd_secret = 357
	__x64_sys_process_mrelease = 358
	__x64_sys_futex_waitv = 359
	__x64_sys_set_mempolicy_home_node = 360
	__x64_sys_cachestat = 361
	__x64_sys_fchmodat2 = 362
	__x64_sys_map_shadow_stack = 363
	__x64_sys_futex_wake = 364
	__x64_sys_futex_wait = 365
	__x64_sys_futex_requeue = 366
	__x64_sys_statmount = 367
	__x64_sys_listmount = 368
	__x64_sys_lsm_get_self_attr = 369
	__x64_sys_lsm_set_self_attr = 370
	__x64_sys_lsm_list_modules = 371
	__x64_sys_mseal = 372
	__x64_sys_setxattrat = 373
	__x64_sys_getxattrat = 374
	__x64_sys_listxattrat = 375
	__x64_sys_removexattrat = 376
	__x64_sys_open_tree_attr = 377
	__x64_sys_file_getattr = 378
	__x64_sys_file_setattr = 379
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

func Mknod(filename []byte, mode uint, dev uint) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_mknod), uintptr(unsafe.Pointer(&filename[0])), uintptr(mode), uintptr(dev))
	return r1, err
}

func Statfs(pathname []byte, buf unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_statfs), uintptr(unsafe.Pointer(&pathname[0])), uintptr(buf), 0)
	return r1, err
}

func SchedSetscheduler(pid int, policy int, param unsafe.Pointer) (uintptr, syscall.Errno) {
	r1, _, err := syscall.Syscall(uintptr(__x64_sys_sched_setscheduler), uintptr(pid), uintptr(policy), uintptr(param))
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
	r1, _, err := syscall.Syscall6(uintptr(__x64_sys_faccessat2), uintptr(dfd), uintptr(unsafe.Pointer(&filename[0])), uintptr(mode), uintptr(flags), 0, 0)
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
