# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

Before any major/minor/patch bump all unit tests will be run to verify they pass.

## [Unreleased]

-   [x]

## [1.2.0] - 2022-07-10

### Added

-   docstrings added to types, methods and functions
-   version retractions added to go.mod

### Changed

-   Entry method renamed from GetRemote to NewRemote
-   Readme updated to reflect latest changes

## [1.1.0] - 2022-06-30

### Added

-   Level updates implemented in Pooler struct. Runs in its own goroutine.

### Fixed

-   Fixed bug with identifier in outputs struct.

### Changed

-   Package files moved into root of repository.
-   Remote struct now exported type

## [1.0.0] - 2022-06-30

### Added

-   recorder, device structs implemented
-   gainlayers field in strip struct implemented
-   levels field in strip, bus structs implemented
-   pooler ratelimit set at 33ms

## [0.0.3] - 2022-06-25

### Added

-   pre-commit.ps1 added for use with git hook
-   unit tests for factory functions added
-   vban parameter methods added
-   support for observers added. publisher/observer structs defined
-   Pooler struct added, pdirty, mdirty now updated continously in a goroutine

### Changed

-   NewRemote factory method now uses director, builder types to create Remote types.
-   cdll renamed to path
-   test suite now using testify/assert

## [0.0.2] - 2022-06-23

### Added

-   physicalStrip, virtualStrip, physicalBus and virtualBus types defined.
-   factory methods for strip, bus now cast return values to interface types.
-   parameter methods added to strip, bus types.
-   command struct implemented
-   bus, vban unit tests added

### Changed

-   strip, bus slices in remote type defined as interface slice types.
-   bindings in base now prepended with vm.
-   vban fields added to kind structs

## [0.0.1] - 2022-06-22

### Added

-   interface entry point defined in remote
-   some base functions are exported through forwarding methods in Remote type (Login, Logout etc)
-   wrapper around the CAPI defined in base
-   path helper functions defined in cdll
-   kind structs defined in kinds. These describe the layout for each version.
-   channel, strip, bus structs getter/setter procedures defined.
-   button struct fully implemented.
-   initial test commit
