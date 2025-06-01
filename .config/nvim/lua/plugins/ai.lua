return {
    "supermaven-inc/supermaven-nvim",
    init = function()
        require("supermaven-nvim").setup({
            keymaps = {
                accept_suggestion = "<C-j>",
                clear_suggestion = "<C-]>",
                accept_word = "<C-i>",
            },
            ignore_filetypes = { cpp = true },
            color = {
                suggestion_color = "#cccccc",
                cterm = 244,
            },
            log_level = "info", -- set to "off" to disable logging completely
            disable_inline_completion = false, -- disables inline completion for use with cmp
            disable_keymaps = false -- disables built in keymaps for more manual control
        })
    end,
}
