#g3status

A custom i3status replacement for use with [i3wm](http://i3wm.org)
This is not a generic solution. It is hard coded for my personal
settings. It is for i3wm and ArchLinux running in VMs.


##Requirements
Because of the Unicode characters I use in the status bar please
make sure you install necessary fonts.

I recommend having DejaVu Sans Mono and Icons installed. DejaVu
is available in official ArchLinux repositories and Icons is
available from AUR at https://aur.archlinux.org/packages/ttf-font-icons/

Since there is no binary available for download, please install
Go on your system to compile the source.

##Installation
Please use _go get_ to install g3status.

   go get github.com/axhixh/g3status


Please edit your i3 configuration file to use *g3status* as
the status command and to use the two recommended fonts.

    bar {
        status_command g3status
        font pango:DejaVu Sans Mono, Icons 9
    }

##Alternatives
If you are interested in more generic Go alternatives to
i3status please look at some of the options at GitHub.

 1. https://github.com/ghedamat/go-i3status
 2. https://github.com/damog/mastodon
