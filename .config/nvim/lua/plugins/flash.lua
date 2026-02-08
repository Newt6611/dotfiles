vim.pack.add({
	{ src = 'https://github.com/folke/flash.nvim' }
})

local fish = require('flash')
fish.setup()

local map = vim.keymap.set


map({ "n", "x", "o" }, "s", fish.jump, { desc = "Flash" })
