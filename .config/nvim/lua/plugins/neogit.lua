return {
    "NeogitOrg/neogit",
    dependencies = {
        "nvim-lua/plenary.nvim",         -- required
        "sindrets/diffview.nvim",        -- optional - Diff integration

        -- Only one of these is needed.
        "nvim-telescope/telescope.nvim", -- optional
        "ibhagwan/fzf-lua",              -- optional
        "echasnovski/mini.pick",         -- optional
        "folke/snacks.nvim",             -- optional
    },

    opts = {
        floating = {
            relative = "editor",
            width = 0.9,
            height = 0.9,
            style = "minimal",
            border = "rounded",
        },
    },

    keys = {
        { "<leader>gg", "<cmd>Neogit kind=floating<cr>" },
    },
}
