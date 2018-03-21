This attempt was to try and recreate an issue we saw maybe 6 months ago with struct with large properties resulting in a memory leak with PyPy.

Unable to recreate at the moment, seems to allocate and de-allocate memory just fine- example spits out 10MB structs, Python consumes one and immediately deletes it over and over and seems fine.