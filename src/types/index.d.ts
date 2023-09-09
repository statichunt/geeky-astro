// home banner types defination
interface Banner {
  title: string;
  title_small: string;
  content: string;
  image_enable: boolean;
  image: string;
  button: Button;
}

//button types definition
export type Button = {
  enable: boolean;
  label: string;
  link: string;
};

// homepage type definition
export interface HomePage {
  banner: Banner;
  featured_posts: {
    enable: boolean;
    title: string;
  };
  promotion: {
    enable: boolean;
    image: string;
    link: string;
  };
  recent_posts: {
    enable: boolean;
    title: string;
  };
}
