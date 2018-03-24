The purpose of this attempt is to work around the multiprocessing limitations encountered in attempt4.

It seems if you import a gopy-built module in a child process, another Go runtime is provisioned and there are no errors.