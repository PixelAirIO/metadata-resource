# Metadata Resource

Fetches the [Build
Metadata](https://concourse-ci.org/docs/resource-types/implementing/#metadata)
for a Concourse job, making it available for later steps in a job.

Add the resource to your pipeline:
```yaml
resource_types:
  - name: metadata
    type: registry-image
    source:
      repository: docker.io/pixelairio/metadata-resource

resources:
  - name: metadata
    type: metadata
```

Use it in your jobs:
```yaml
jobs:
  - name: my-job
    plan:
      - put: metadata
      - load_var: metadata
        file: metadata/build.json
        reveal: true
      # Can then reference using the local var
      # ((.:metadata.build_url))
```

The published image is a minimal container image that contains only the
resource binaries. No shell, package manager, or unix tooling are included.

## `source` Configuration

This resource has no configuration and implements a no-op for its check. This
resource should not be used to trigger jobs.

## Put Step

The put step emits a version that ensures no resource cache is used. **Do not
set `no_get` to `true`**. This resource relies on the implicit `get` step
running in order to make the build metadata available in later steps.

This step has no params.

## Get Step

Do not directly use the get step for this resource. Use the put step.

This step has no params.

This step emits the Build Metadata in a `build.json` file in the following format. All fields will be present even if they have no value:

```json
{
  "build_id": "456", // Globally unique ID that Concourse uses in the backend to identify the job
  "build_name": "3", // Build number within the scope of the job. The one you see in the web UI
  "job_name": "",
  "pipeline_name": "",
  "instance_vars": {}, // Will be a map of the pipeline's instance vars
  "team_name": "",
  "created_by": "",
  "external_url": "",
  "build_url": "",
  "build_url_short": "",
}
```

## Development

Build the image

```
docker build metadata-resource .
```
