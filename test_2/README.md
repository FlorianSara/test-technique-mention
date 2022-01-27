## Cache Library

Write a Cache library with support for multiple backends.

The library must allow users to do two operations:

 - `set($key, $value)`: sets the item named `$key` in the cache, to value `$value`
 - `get($key)`: returns the value stored in the cache for item `$key`, or `NULL`
   if it doesn't exist

`$key` is always a string. `$value` can be anything.

The library must have two backends:

 - File: Stores values in local files.
 - Mem: Stores values in the memory of the current process.

The library doesn't have an opinion on which backend to use.

Language is up to you.

The answer will be rated based on object-oriented conception, simplicity, and reliability.

### Some context

In software, a cache is a simple storage that can be used to store data for later. For example, a web page could store the result of a SQL query in a cache: Upon future requests, the web page can avoid the SQL query execution if the value is already in the cache, thus saving some time.

A cache is most useful when its data outlives the process or request that set it, because the data will be available to future or concurrent processes or requests. Less persistent backends like the "Mem" one can still be useful when a process or request is caching data for itself or when caching is not necessary in some environment.