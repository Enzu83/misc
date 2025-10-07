use std::ffi::{c_char, CStr};

static WORLD_STR: &[u8] = b"world\0";

#[unsafe(no_mangle)]
extern "C" fn ffi_func(cstr_ptr_ptr: *mut *mut c_char) {
    let new_cstr = CStr::from_bytes_with_nul(WORLD_STR).unwrap();

    unsafe { *cstr_ptr_ptr = new_cstr.as_ptr() as *mut c_char; }
}
