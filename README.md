# Backend

### Structure
  - main.go
  - config/
  - handlers/
  - models/
  - repositories/
  - services/
  - middlewares/
  - utils/
    
#### `main.go`
This file serves as the entry point of the application. It typically contains the main function that starts the application and sets up any necessary configurations or dependencies.

#### `config/` 
This directory is used to store constants, configuration files, or environment variables that the application requires. It helps separate configuration-related code from the rest of the application logic.

#### `handlers/`
The handlers directory typically contains HTTP request/response handlers. These handlers handle incoming HTTP requests, process them, and generate appropriate responses. They are responsible for handling the HTTP communication between the client and the server.

#### `models/`
The models directory contains data models or structs that represent the application's data structures or entities. These models define the structure of the data and may include validation rules or methods associated with the data.

#### `repositories/`
The repositories directory typically contains code related to database operations or interactions with third-party services. It encapsulates the logic for reading from and writing to databases or other external systems.

#### `services/`
The services directory houses the implementation of the application's business logic. It contains code that performs operations using the data models and repositories. Services typically encapsulate the core functionality of the application.

#### `middlewares/`
The middlewares directory is used for implementing interceptors or middleware functions. These functions can intercept and modify incoming requests or outgoing responses, perform additional processing, or add functionality such as logging, authentication, or request/response validation.

#### `utils/`
The utils directory typically contains helper functions or utility code that is shared across different parts of the application. It provides reusable functions that assist with common tasks or provide additional functionality.

### Points to note

##### 1) Directory Structure
Each module, service, or package should have its own separate directory within the specified set. This helps maintain a modular and organized codebase.
Follow a consistent naming convention for directories and use meaningful names that reflect the purpose of the module, service, or package.
Avoid having too many nested directories to keep the structure simple and readable.

##### 2) Exporting Functions and Models
Properly export functions and models by starting their names with an uppercase letter if they need to be accessible outside of their respective packages.
Use appropriate structuring and package design to avoid circular dependency issues. Plan the dependencies between packages carefully to prevent circular references.

##### 3) Configuration Files
Configuration files should not be updated without confirmation from any of the lead members. Changes to configuration files can have a significant impact on the application's behavior, so it is important to ensure that modifications are reviewed and approved by relevant team members or leads.
Consider using a version control system to track changes to configuration files and to enable collaboration and review processes
