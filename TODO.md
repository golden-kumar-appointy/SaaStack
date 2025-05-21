Step 1: 

- Register 1 Interface
- Register 1 Plugin that implements the interface
- Make an HTTP API call to the platform which should resolve in the plugin
    - Core - Doesn’t Import
    - Interface - Can Import core only
    - Plugin - Can Import core and interface
    - Main - Can imports all three