-- Pull in the wezterm API
local wezterm = require 'wezterm'

-- This will hold the configuration.
local config = wezterm.config_builder()

config.default_prog = { '/bin/zsh', '-l', '-c', 'zellij -l welcome' }

config.max_fps = 120
config.window_decorations = 'RESIZE'
config.color_scheme = 'Batman'

config.font = wezterm.font 'JetBrainsMono Nerd Font Mono'
config.font_size = 24

config.enable_tab_bar = false
config.window_background_opacity = 1
config.color_scheme = 'Sex Colors (terminal.sexy)'

-- config.term = 'wezterm'

-- and finally, return the configuration to wezterm
return config
