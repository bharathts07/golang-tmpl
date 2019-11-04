<a name="unreleased"></a>
## [Unreleased]


<a name="initial"></a>
## initial - 2019-11-04
### Add
- **.circleci:** config file
- **.github:** issue and PR templates
- **Makefile:** with simple commands
- **Procfile:** for heroku execution
- **README:** for clarity
- **assets:** to serve static artifacts
- **config:** pkg for parsing environment values
- **database:** template interface declaration
- **di:** pkg for dependency injection
- **docker:** make command and impl for building docker image
- **fakedb:** implementation for database.Client
- **global:** pkg
- **handlers:** jokes endpoint handler functions
- **http:** env config at server start
- **http:** route grouping and versioning
- **internal:** models pkg for storing internal data structs
- **joke:** service and handler implementation
- **log:** pkg
- **middleware:** for storing and extracting db from gin context
- **pokke:** graceful shutdown to the server
- **server:** template server that uses gin-gonic pkg
- **transport:** function comments

### Fix
- **pokke:** lint
- **pokke:** lint

### Refactor
- **database:** remove all fakeDB usages
- **http:** imports
- **pokke:** fix invalid pointer refernce
- **server:** seperate under transport root package
- **server:** remove graceful shutdown
- **server:** handlers and routes to differnt pkgs
- **server:** main logic to a seperate func

### Regenerate
- **database:** mock

### Remove
- **handler:** home and layout handlers

### Rename
- **models:** HomeLayout-> Layout

### Update
- **.gitignore:** add additional folders to be ignored
- **Makefile:** add make dependecies and coverage cmds
- **Makefile:** add mockgenerate command for database interface
- **Makefile:** add command to push docker image to hub
- **circleci:** remove org-global context variable
- **cmd:** to use the updated server pkg
- **cmd:** create a fakedb client for use in server
- **config:** add additional logging for clarity
- **database:** interface function signature
- **env:** struct fields
- **global:** add ContextDBName for use as key in gin Context
- **go-mod:** dependencies
- **handlers:** GetWelcomeComponents() to extract and use db
- **internal:** Component fields data type signature
- **pokke:** add best_practises source link
- **pokke:** read server port from env
- **pokke:** read only PORT
- **pokke:** call a GetPort()
- **server:** newlines at end of fmt logs
- **server:** db from middleware in route handlers


[Unreleased]: https://github.com/bharathts07/pokke/compare/initial...HEAD
