# Liferay Build Action

`lgdd/liferay-build-action@v1` build your Liferay workspace (Gradle or Maven) and can upload the bundle and artifacts in the workflow to be used later (e.g. for a release).

## Usage

You can use this action in a [GitHub Actions Workflow](https://help.github.com/en/articles/about-github-actions) by a adding a YAML file under `.github/workflows/` with the following content:

### Examples

#### Build on push & pull request

```yaml
name: liferay-workspace-build
run-name: Liferay Workspace Build

on:
  push:
  pull_request:

jobs:
  liferay-build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - uses: lgdd/liferay-build-action@v1
        with:
          profile: 'prod'
          upload-bundle: false
          upload-artifacts: false
```

#### Release with bundle only

```yaml
name: liferay-workspace-release
run-name: Liferay Workspace Release

on:
  push:
    tags:
      - 'v*'

jobs:
  liferay-build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - uses: lgdd/liferay-build-action@v1
        with:
          profile: 'prod'
          upload-bundle: true
          upload-artifacts: false
      - uses: actions/download-artifact@v3
        with:
          name: bundle
      - name: Create Release
        if: success()
        env:
          GH_TOKEN: ${{ github.token }}
        run: |
          ./gradlew deploy
          gh release create ${{ github.ref }} $(ls *.zip) --generate-notes
```

#### Release with artifacts only

```yaml
name: liferay-workspace-release
run-name: Liferay Workspace Release

on:
  push:
    tags:
      - 'v*'

jobs:
  liferay-build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - uses: lgdd/liferay-build-action@v1
        with:
          profile: 'prod'
          upload-bundle: false
          upload-artifacts: true
      - uses: actions/download-artifact@v3
        with:
          name: artifacts
      - name: Create Release
        if: success()
        env:
          GH_TOKEN: ${{ github.token }}
        run: |
          ./gradlew deploy
          gh release create ${{ github.ref }} $(ls *.jar && ls *.zip && ls *.war) --generate-notes
```

## License

[MIT](LICENSE)
