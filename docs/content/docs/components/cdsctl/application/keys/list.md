---
title: "list"
notitle: true
notoc: true
---
# cdsctl application keys list

`List CDS application keys`

## Synopsis

`List CDS application keys`

```
cdsctl application keys list [ PROJECT-KEY APPLICATION-NAME ] [flags]
```

## Options

```
      --fields string   Only display specified object fields. 'empty' will display all fields, 'all' will display all object fields, 'field1,field2' to select multiple fields
      --filter string   Filter output based on conditions provided
      --format string   Output format: table|json|yaml (default "table")
  -q, --quiet           Only display object's key
```

## Options inherited from parent commands

```
  -f, --file string   set configuration file
  -k, --insecure      (SSL) This option explicitly allows curl to perform "insecure" SSL connections and transfers.
  -w, --no-warnings   do not display warnings
  -v, --verbose       verbose output
```

## SEE ALSO

* [cdsctl application keys](/docs/components/cdsctl/application/keys/)	 - `Manage CDS application keys`

