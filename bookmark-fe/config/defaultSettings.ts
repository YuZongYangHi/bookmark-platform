import { ProLayoutProps } from '@ant-design/pro-components';

/**
 * @name
 */
const Settings: ProLayoutProps & {
  pwa?: boolean;
  logo?: string;
} = {
  "navTheme": "light",
  "colorPrimary": "#52C41A",
  "layout": "side",
  "contentWidth": "Fluid",
  "fixedHeader": false,
  "fixSiderbar": true,
  "pwa": true,
  "menu": {
    locale: false
  },
  "logo": "/img/bookmarks.png",
  "title": "BookMark",
  "splitMenus": false,
  "siderMenuType": "sub",
  "token": {
    "sider": {
      colorMenuBackground: '#fff',
      colorMenuItemDivider: '#dfdfdf',
      colorTextMenu: '#595959',
      colorTextMenuSelected: '#52C41A',
      colorBgMenuItemSelected: 'rgba(230,243,254,1)',
    }
  }
}

export default Settings;
