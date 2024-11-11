export interface SidebarItem {
	text: string;
	href?: string;
	subitems?: { text: string; href: string }[];
}
