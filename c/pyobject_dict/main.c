#include <Python.h>

void add_new_key(PyObject *dict, char *key_str, char *value_str) {
    // Add some key-value pairs to demonstrate it works
    PyObject *key = PyUnicode_FromString(key_str);
    PyObject *value = PyUnicode_FromString(value_str);
    
        
    if (key && value) {
        PyDict_SetItem(dict, key, value);
        printf("Added key-value '%s: %s' pair to dictionary\n", key_str, value_str);
    }

    Py_XDECREF(key);
    Py_XDECREF(value);
}

int main() {
    // Initialize the Python interpreter
    Py_Initialize();
    
    // Create a new dictionary PyObject
    PyObject *dict = PyDict_New();
    
    if (dict == NULL) {
        fprintf(stderr, "Failed to create dictionary\n");
        Py_Finalize();
        return 1;
    }
    
    // Verify that it's actually a dictionary
    if (PyDict_Check(dict)) {
        printf("Successfully created a Python dictionary object\n");
    } else {
        printf("Failed: Object is not a dictionary\n");
    }
    
    add_new_key(dict, "name", "My first service");
    add_new_key(dict, "url", "http://datadoghq.com");
    
    // Clean up
    Py_DECREF(dict);
    
    // Finalize the Python interpreter
    Py_Finalize();
    
    return 0;
}
