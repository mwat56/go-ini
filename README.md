# INI

## Purpose

Over the times several differente file formats have been developed just for storing configuration data for some program.
While they all may have some merits for me the two-dimensional INI file format – made popular by the first MS-Windows versions – was always sufficient for my needs.
This package provides the `TSections` class to read/parse, modify, and write such INI files.

## Installation

You can use `Go` to install this package for you:

    go get github.com/mwat56/go-ini

## Usage

An INI file usually looks like this:

    # This is a comment

    [aSectionName]
    key1 = value 1
    key2 = value2
    …

    [anotherSection]
    key1 = value1
    key2 = value 2
    …

You can create a `TSections` instance by either calling `ini.NewSections()` and then using the numerous methods (including `Load()` and `Store()`).
Or you simply call `ini.LoadFile(aFilename)` which does – as the name suggests – the loading for you.

## Licence

    Copyright (C) 2019  M.Watermann, 10247 Berlin, FRG
                    All rights reserved
                EMail : <support@mwat.de>

> This program is free software; you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation; either version 3 of the License, or (at your option) any later version.
>
> This software is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
>
> You should have received a copy of the GNU General Public License along with this program.  If not, see the [GNU General Public License](http://www.gnu.org/licenses/gpl.html) for details.