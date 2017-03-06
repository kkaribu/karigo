# karigo

karigo is an API framework based on [jsonapi](https://github.com/kkaribu/jsonapi).

Here's a list of features supported by the framework:

 * Accept JSON API requests and return JSON API responses ([jsonapi.org/format](http://jsonapi.org/format))
 * Manipulate resources, collections, and relationships
 * The Store interface can be implemented to use any DBMS (implementation for PostgreSQL already provided)
 * Automatically synchronize the schema of your database so that it matches the types you define
 * Utilities to make your app a command line application
  * If you make an app called myapp, you can run it with the command `myapp run`, not `karigo run myapp`
  * You also get other tools like `myapp sync` to update your database's schema
  * Those commands are meant to be run in a directory that contains a file with the necessary configuration
 * Use JWT for sessions ([jwt.io](https://jwt.io))

## State

The framework has no release nor documentation yet.
