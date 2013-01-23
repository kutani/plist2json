plist2json - a no-frills converter in golang
======

Currently in a basic, works-for-me state.  I wrote it when I was
unable to find any decent non-OS X tools for doing this conversion.

At current takes a plist file as an argument and outputs JSON to
stdout.  If you want pretty printing, pipe it to a pretty printer.

Does not support the full plist spec yet; will if I get a patch or
if I find a need, or if I get bored enough to do it.

See LICENSE for licensing information
