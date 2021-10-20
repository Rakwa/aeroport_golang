import { defineConfig } from 'windicss/helpers'
import colors from 'windicss/colors'

export default defineConfig({
  theme: {
    extend: {
      fontFamily: {
        'oxygen': ['"Oxygen"']
      },
      screens: {
        sm: '640px',
        md: '768px',
        lg: '1024px',
        xl: '1280px',
        '2xl': '1536px',
      },
      colors: {
        primary: '#0096C7',
        primaryBg: '#007997',
        // ...
      },
      
    },
  },
})
