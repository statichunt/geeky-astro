export interface Schema {
	categories: Category[];
  posts: Post[];
  landing_pages: LandingPage[];
}

export interface Posts {
  posts: Post[];
}

export interface Category {
  category: string;
  url: string;
}

export interface PostName {
  slug: string;
}

export interface Post {
  id: number;
  sort: number;
  title: string;
  slug: string;
  tags: string[];
  featured: boolean;
  image: string;
  body: string;
  status: string;
  date_created: string;
}

export interface LandingPage {
  id: number;
  name: string;
  enter_msisdn: boolean;
  enter_pin: boolean;
  background_color: string;
  cta_text: string;
  image: string;
  image_full_screen: boolean;
  button_color: string;
  button_text_color: string;
  button_text: string;
  navigate_to: string;
  bottom_text: string;
  footer: string;
  status: string;
}
