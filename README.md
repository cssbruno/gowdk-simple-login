# GOWDK Simple Login

 Minimal login app built with GOWDK.

## Run

```sh
gowdk build --target gowdk-simple-login
GOWDK_ADDR=127.0.0.1:8090 bin/gowdk-simple-login
```

Open `http://127.0.0.1:8090/`.

Credentials:

```text
email: demo@example.com
password: demo-password
```

## VS Code

Open the project workspace file so GOWDK diagnostics run from this project
root:

```sh
code gowdk-simple-login.code-workspace
```

If VS Code is opened from a parent folder instead of this project root, the
extension can run `gowdk check` from the wrong directory and report
`no .gwdk files found`.

## Files

- `src/auth/login.page.gwdk`: login page, `Login` action, and `Session` API declaration.
- `src/auth/password-hint.cmp.gwdk`: reactive password field with `client {}` state, computed values, binding, and events.
- `src/auth/dashboard.page.gwdk`: signed-in dashboard and `Logout` action.
- `src/auth/auth.go`: normal Go auth handlers.
- `styles/app.css`: page stylesheet.
- `gowdk.config.go`: binary target config. The target output is inferred as
  `.gowdk/output/gowdk-simple-login`, generated app source is written to
  `.gowdk/gowdk-simple-login`, and the binary is `bin/gowdk-simple-login`.

## Check

```sh
gowdk check src/auth/*.gwdk
go test ./...
```
