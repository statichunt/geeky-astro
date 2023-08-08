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

// aboutpage type definition
export interface AboutPage {
  title: string;
  image: string;
  description: string;
  education: {
    title: string;
    degrees: Array<{
      university: string;
      content: string;
    }>;
  };
  experience: {
    title: string;
    list: Array<string>;
  };
}

// contact page type definition
export interface ContactPage {
  title: string;
  description: string;
  form: string;
  social_links: Array<{
    icon: string;
    content: string;
  }>;
}
