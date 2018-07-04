## Contributing Guidelines

Thank you for considering to contribute to this library!

The following text lists guidelines for contributions.
These guidelines don't have legal status, so use them as a reference and common sense - and feel free to update them as well!


### "I just want to know..."

For questions, or general usage-information about **Dear ImGui**, please refer to the [homepage](https://github.com/ocornut/imgui), or look in the detailed documentation of the C++ source.
This wrapper houses next to no functionality. As long as it is not a Go-specific issue, help will rather be there.

### Scope

This wrapper exposes minimal functionality of **Dear ImGui**. Ideally, this functionality is that of the common minimum that someone would want. This wrapper does not strife for full configurability, like the original library. This is not even possible in some cases, as it requires compilation flags.

### Extensions
At the moment, this library is primarily used by **InkyBlackness**. If you can and want to make use of this library in your own projects, you are happy to do so. Pull-requests with extensions are happily accepted, provided that they uphold the following minimum requirements:
* Code is properly formatted & linted (use [metalinter](https://github.com/alecthomas/gometalinter) for a full check)
* Public Go API is documented (copied documentation from **Dear ImGui** is acceptable and recommended, assuming it is adapted regarding type names)
* API and version philosophies are respected (see README.md)
