# Changelog

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

- Add more documentation about how RegisterPreserveFields works so people
  don't have to chase the godocs to figure out how to implement it.

## [2.0.0-rc5] - 2018-07-04

### Changed

- The upstream golang.org/x/oauth2 library has changed it's API, this fixes
  the breakage.

## [2.0.0-rc4] - 2018-06-27

### Changed

- RememberingServerStorer now has context on its methods

## [2.0.0-rc3] - 2018-05-25

### Changed

- Recover and Confirm now use split tokens

    The reason for this change is that there's a timing attack possible
    because of the use of memcmp() by databases to check if the token exists.
    By using a separate piece of the token as a selector, we use memcmp() in
    one place, but a crypto constant time compare in the other to check the
    other value, and this value cannot be leaked by timing, and since you need
    both to recover/confirm as the user, this attack should now be mitigated.

    This requires users to implement additional fields on the user and rename
    the Storer methods.

## [2.0.0-rc2] - 2018-05-14

Mostly rewrote Authboss by changing many of the core interfaces. This release
is instrumental in providing better support for integrating with many web frameworks
and setups.

### Added

- v2 Upgrade guide (tov2.md)

- API/JSON Support

    Because of the new abstractions it's possible to implement body readers,
    responders, redirectors and renderers that all speak JSON (or anything else for that
    matter). There are a number of these that exist already in the defaults package.

### Changed

- The core functionality of authboss is now delivered over a set of interfaces

    This change was fairly massive. We've abstracted the HTTP stack completely
    so that authboss isn't really doing things like issuing template renderings,
    it's just asking a small interface to do it instead. The reason for doing this
    was because the previous design was too inflexible and wouldn't integrate nicely
    with various frameworks etc. The defaults package helps fill in the gaps for typical
    use cases.

- Storage is now done by many small interfaces

    It became apparent than the old reflect-based mapping was a horrible solution
    to passing data back and forth between these structs. So instead we've created a
    much more verbose (but type safe) set of interfaces to govern which fields we need.

    Now we can check that our structs have the correct methods using variable declarations
    and there's no more confusion about how various types map back and forth inside the
    mystical `Bind` and `Unbind` methods.

    The downside to this of course is it's incredibly verbose to create a fully featured
    model, but I think that the benefits outweigh the downsides (see bugs in the past about
    different types being broken/not supported/not working correctly).

- Support for context.Context is now much better

    We had a few pull requests that kind of shoved context.Context support in the sides
    so that authboss would work in Google App Engine. With this release context is
    almost everywhere that an external system would be interacted with.

- Client State management rewritten

    The old method of client state management performed writes too frequently. By using a
    collection of state change events that are later applied in a single write operation at
    the end, we make it so we don't get duplicate cookies etc. The bad thing about this is
    that we have to wrap the ResponseWriter. But there's an UnderlyingResponseWriter
    interface to deal with this problem.

- Validation has been broken into smaller and hopefully nicer interfaces

    Validation needs to be handled by the BodyReader's set of returned structs. This punts
    validation outside of the realm of Authboss for the most part, but there's still
    helpful tools in the defaults package to help with validation if you're against writing
    rolling your own.

- Logout has been broken out into it's own module to avoid duplication inside login/oauth2
  since they perform the same function.

- Config is now a nested struct, this helps organize the properties a little better (but
  I hope you never mouse over the type definition in a code editor).

### Removed

- Notable removal of AllowInsecureLoginAfterConfirm

### Fixed

- Fix bug where e-mail with only a textbody would send blank e-mails

### Deprecated

- Use of gopkg.in, it's no longer a supported method of consuming authboss. Use
  manual vendoring, dep or vgo.

## [1.0.0] - 2015-08-02
### Changed
This change is potentially breaking, it did break the sample since the supporting struct was wrong for the data we were using.

**Lock:** The documentation was updated to reflect that the struct value for AttemptNumber is indeed an int64.
**Unbind:** Previously it would scrape the struct for the supported types (string, int, bool, time.Time, sql.Scanner/driver.Valuer)
and make them into a map. Now the field list will contain all types found in the struct.
**Bind:** Before this would only set the supported types (described above), now it attempts to set all values. It does check to ensure
the type in the attribute map matches what's in the struct before assignment.

## 2015-04-01 Refactor for Multi-tenancy
### Changed
This breaking change allows multiple sites running off the same code base to each use different configurations of Authboss. To migrate
your code simply use authboss.New() to get an instance of Authboss and all the old things that used to be in the authboss package are
now there. See [this commit to the sample](https://github.com/volatiletech/authboss-sample/commit/eea55fc3b03855d4e9fb63577d72ce8ff0cd4079)
to see precisely how to make these changes.
