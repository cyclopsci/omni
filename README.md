omni
====

Omni is a command line utility to help manage multiple automation platforms installed on a single machine.  It provides a thin wrapper around bundler and virtualenv to give a consistent install and exec interface into multiple puppet and ansible execution environments.

## Dependencies
For ruby-based platforms (puppet):
 * ruby
 * bundler

For python-based platforms (ansible):
 * python
 * pip
 * virtualenv

## Usage
```
# Install a new platform version
omni install puppet 4.0.0

# Enter the version and execute commands in your shell
$(omni enter puppet 4.0.0)
puppet --version
$(omni exit)

# Or skip entering and run an exec
omni exec -p puppet -v 4.0.0 -- pupet --version        # Note the -- to stop flag parsing
omni exec -p puppet -v 3.7.5,4.0.0 -- puppet --version # Run a command against multiple versions of puppet
