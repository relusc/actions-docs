# Usage examples

All examples are based on this folder's [`action.yml`](./action.yml).

## Synopsis of `markdown/md` command

```text
actions-docs markdown [command options]
```

## Options

```text
--create-file   (default: false)
--format value  (default: "table")
--help, -h      show help (default: false)
```

## Intputs and outputs as table printed to stdout

```sh
actions-docs md
```

```text
## Inputs

| Name | Description | Required | Default |
|------|-------------|----------|:-------:|
| `greeting` | Which greeting do you want to use | no | `Howdy` |
| `version` | Version of the action | no | `v1.2.3` |
| `who-to-greet` | Who do you want to greet | yes | - |

## Outputs

| Name | Description |
|------|:-----------:|
| who-did-i-greet | Who did you greet |
```

## Inputs and outputs as list printed to stdout

```sh
actions-docs md --format=list
```

```text
## Inputs

- `greeting`: Which greeting do you want to use. Is not required. Defaults to `Howdy`
- `version`: Version of the action. Is not required. Defaults to `v1.2.3`
- `who-to-greet`: Who do you want to greet. Is required.

## Outputs

- `who-did-i-greet`: Who did you greet
```

## Generate file with inputs and outputs as list

```sh
actions-docs md --create-file --format=list
```

Creates a `README.md` file with the content:

```text
# example-action

Sample action which basically does nothing.

## Inputs

- `greeting`: Which greeting do you want to use. Is not required. Defaults to `Howdy`
- `version`: Version of the action. Is not required. Defaults to `v1.2.3`
- `who-to-greet`: Who do you want to greet. Is required.

## Outputs

- `who-did-i-greet`: Who did you greet
```
