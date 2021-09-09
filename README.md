# terraform-provider-toolbox

A magical box of tools for your Terraform recipes.

## Building the provider

Run the following command to build the provider

```bash
$> make build
```

## Installing the provider locally

Run the following command to install the provider locally:

```bash
$> make install
```

## Test sample configuration

First, build and install the provider.

```bash
$> make install
```

Then, move into the `examples` directory and run the following command to initialize the workspace and apply the sample configuration.

```bash
$> terraform init && terraform apply
```

If you rebuild the provider, you will need to remove the local `.terraform/` directory.

## Debugging the provider

Run the following:

```bash
$> export TF_LOG_CORE=off
$> export TF_LOG_PROVIDER=TRACE
$> cd examples/
$> cd .. && make install && cd examples && terraform init && terraform apply
```
