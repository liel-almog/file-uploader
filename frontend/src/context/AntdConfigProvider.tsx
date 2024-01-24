import { ConfigProvider } from "antd";
import { PropsWithChildren } from "react";
import locale from "antd/locale/he_IL";

export const AntdConfigProvider = ({ children }: PropsWithChildren) => (
  <ConfigProvider
    theme={{
      token: {
        colorError: "hsla(0, 84%, 64%, 1)",
        blue: "#4c65b8",
        fontFamily: "Noto Sans Hebrew",
      },
      cssVar: true,
    }}
    direction="rtl"
    locale={locale}
  >
    {children}
  </ConfigProvider>
);
