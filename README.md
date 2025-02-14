# Liferay Build Action

`lgdd/liferay-build-action@v2` build your Liferay workspace (Gradle or Maven) and can upload the bundle and artifacts in the workflow to be used later (e.g. for a release).

## v2

### New
- **build-extra-args**: this new parameter allows you to pass extra arguments to the build step of this action. For example, if you want to skip tests with Gradle you can now add `build-extra-args: '-x test'`, and if you're using Maven you would add `build-extra-args: '-DskipTests=true'`.

### Breaking Changes
#### GitHub Actions Artifact Changes
In v1, this action was using the `actions/upload-artifact@v3` and `actions/download-artifact@v3` which are [deprecated](https://github.blog/changelog/2024-04-16-deprecation-notice-v3-of-the-artifact-actions/). One of the impact here is the change is the name of the artifact uploaded. Now it follows the pattern `bundle--jdk-${java-version}` for the Liferay Bundle, and `artifacts--jdk-${java-version}` for the Liferay Modules/Client Extensions.

Here's an example of changes you would need to make while upgrading to `actions/download-artifact@v4` and `lgdd/liferay-build-action@v2`:

```diff
steps:
   - uses: actions/checkout@v4
   - uses: lgdd/liferay-build-action@v2
     with:
       java-version: 21
       upload-bundle: true
       upload-artifacts: true
-    - uses: actions/download-artifact@v3
+    - uses: actions/download-artifact@v4
       with:
-        name: bundle
+        name: bundle--jdk-21
-    - uses: actions/download-artifact@v3
+    - uses: actions/download-artifact@v4
       with:
-        name: artifacts
+        name: artifacts--jdk-21
```

#### Default Java Version

Since Liferay DXP DXP 2024.Q2 and Liferay Portal 7.4 GA120 are the last version to support Java 11, the v2 of this action is using Java 17 as the new default Java version to align with Liferay changes.

If you still need to build projects with Java 11 because your Liferay version is prior to the one mentioned before, you should explicitly add it as follows:

```diff
steps:
   - uses: actions/checkout@v4
   - uses: lgdd/liferay-build-action@v2
     with:
+      java-version: 11
```

## Usage

You can use this action in a [GitHub Actions Workflow](https://help.github.com/en/articles/about-github-actions) by a adding a YAML file under `.github/workflows/` with the following content:

### Examples

#### Build on push & pull request

```yaml
name: Liferay Workspace Build

on:
  push:
  pull_request:

jobs:
  liferay-build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: lgdd/liferay-build-action@v2
        with:
          profile: 'prod'
          upload-bundle: false
          upload-artifacts: false
```

#### Release with bundle only

```yaml
name: Liferay Workspace Release

on:
  push:
    tags:
      - 'v*'

jobs:
  liferay-build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: lgdd/liferay-build-action@v2
        with:
          profile: 'prod'
          upload-bundle: true
          upload-artifacts: false
      - uses: actions/download-artifact@v4
        with:
          name: bundle--jdk-17
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
name: Liferay Workspace Release

on:
  push:
    tags:
      - 'v*'

jobs:
  liferay-build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: lgdd/liferay-build-action@v2
        with:
          profile: 'prod'
          upload-bundle: false
          upload-artifacts: true
      - uses: actions/download-artifact@v4
        with:
          name: artifacts--jdk-17
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
