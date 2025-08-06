#[repr(C)]
pub struct Instance {
    pub cb: Option<extern "C" fn()>,
}

#[unsafe(no_mangle)]
pub extern "C" fn Run(instance_ptr: *const Instance) {
    if instance_ptr.is_null() {
        println!("NULL instance received");
    }

    let instance = unsafe { &*instance_ptr };

    if let Some(cb) = instance.cb {
        cb();
    } else {
        println!("Received null pointer to exported function!")
    }
}
