# Gnome Go Env

A environment variable module with ehancements arounds
expanding variables, handling environment PATH, and
windows environment variables.

The enhanced `env.Expand()` function handles bash
variables such as `$VAR`, `${VAR}`, `${VAR:-DefaultValue}`
`${VAR:?Error Message}`, and `${VAR:=DefaultValue}`.

The `env.GetWin` and `env.SetWin` methods get and set
environment variables directly from the windows registry.
