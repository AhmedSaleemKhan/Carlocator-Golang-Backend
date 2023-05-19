module.exports = {
  content: ["./public/**/*.html", "./src/**/*.{js,jsx,ts,tsx}"],

  theme: {
    screens: {
      mobile: "340px",

      tab: "780px",

      desktop: "1240px",
    },

    colors: {
      dark_silver: "#37474F",

      input_bg: "#F5F5F5",

      input_text: "#A6A6A6",

      anchor: "#2D98DA",

      submit_button: "#2D98DA",

      white: "#FFFFFF",

      button_hover: "#6773EB",

      blue: "#2D98DA",

      dark_blue: "#0A196F",

      light_blue: "#DCF3FF",

      dull_blue: "#2C87CC",

      header_blue: "#2C87CC",

      time_black: "#292D32",

      light_yellow: "#F5FCD2",

      active_button: "#E0E9F4",

      modal_text: "#1A202C",

      modal_placeholder: "#040815",

      assign_staff: "#D6E4FD",

      reset_button: "#EBEBEB",

      red: "#FF0000",

      dull: "#90A3BF",

      dull_green: "#9CD323",

      dull_black: "#C7C7C7",

      warning: "#FFC73A",
    },
    extend: {
      backgroundImage: {
        bgAuth: "url(../src/images/bgAuth.png)",
      },
    },
  },

  plugins: [],
};
