# ✅ Pros

* Use of a `Makefile` for automation
* Docker Compose setup provided
* Dockerfile
* Validation on requests
* Architecture

# ❌ Cons

* Singleton usage for `MongoClient` (could lead to issues in concurrent scenarios)
* No unit tests included
* No bench performance tests
* Use of `vendor` directory alongside `go.mod` (redundant and discouraged)
* No use of dragonfly
* Errors returned in raw format

# 📝 Conclusion

* The use of the `vendor` directory is generally deprecated in modern Go projects using modules (`go.mod`). It’s better to rely solely on module-based dependency management unless there's a specific requirement.
* While the `Makefile` includes a command for running Docker Compose, it would be helpful to also include targets for building the Go binary or running tests to improve developer experience.
* We can see an influence of other languages/libraries architecture in this project.
Some good Go good practice missing but overall a good project.
