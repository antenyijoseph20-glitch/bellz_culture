package handlers

import (
	"bytes"
	"html/template"
	"io/fs"
	"log"
	"net/http"
)

type Product struct {
	ID          string
	Name        string
	Category    string
	Price       string
	ImageURL    string
	Description string
	Tagline     string
	Badge       string
}

type PageData struct {
	BrandName   string
	Tagline     string
	PhoneNumber string
	Email       string
	Address     string
	HeroImage   string
	CraftImage  string
	FounderBio  string
	Products    []Product
}

// NewHomeHandler parses the embedded homepage and returns its HTTP handler.
func NewHomeHandler(templateFS fs.FS) (http.Handler, error) {
	tmpl, err := template.ParseFS(templateFS, "index.html")
	if err != nil {
		return nil, err
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet && r.Method != http.MethodHead {
			w.Header().Set("Allow", "GET, HEAD")
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		data := PageData{
			BrandName:   "Bellz Culture",
			Tagline:     "Handcrafted Luxury Leather Shoes & Bags",
			PhoneNumber: "2349079997307",
			Email:       "bellzculture@gmail.com",
			Address:     "Oron Road, by Zenith Bank, Uyo, Akwa Ibom State, Nigeria",
			HeroImage:   "/static/images/marketing.jpeg",
			CraftImage:  "/static/images/workshop-2.jpeg",
			FounderBio:  "A young, vibrant, dedicated, and strong-willed artisan who began crafting shoes and bags to support her education. After graduation and marriage, driven by her deep passion and enduring love for the craft, she officially established Bellz Culture as a premier destination for luxury handmade leather goods.",
			Products: []Product{
				{
					ID:          "1",
					Name:        "Signature Striped Utility Tote",
					Category:    "Handbags & Totes",
					Price:       "₦42,000",
					ImageURL:    "/static/images/bags_sample-3.jpeg",
					Description: "Versatile woven striped canvas and leather trim tote built for durability, travel, and daily office elegance.",
					Tagline:     "Spacious style for the active professional.",
					Badge:       "Best Seller",
				},
				{
					ID:          "2",
					Name:        "Structured Mini Lock Bag",
					Category:    "Luxury Handbags",
					Price:       "₦38,000",
					ImageURL:    "/static/images/special-bags.jpeg",
					Description: "Compact box tote featuring woven striped pattern inserts, high-grade leather trim, and gold-tone lock hardware.",
					Tagline:     "Distinctive heritage craft for modern style.",
					Badge:       "New Arrival",
				},
				{
					ID:          "3",
					Name:        "Red Croc-Embossed Statement Bag",
					Category:    "Evening Bags",
					Price:       "₦40,000",
					ImageURL:    "/static/images/special-demand.jpeg",
					Description: "Embossed croc-texture red leather bag with woven gold chain accents and secure front flap closure.",
					Tagline:     "Bold design crafted to stand out.",
					Badge:       "Trending",
				},
				{
					ID:          "4",
					Name:        "Woven Buckle Accent Slides",
					Category:    "Footwear",
					Price:       "₦25,000",
					ImageURL:    "/static/images/foots-wear4.jpeg",
					Description: "Chic open-toe slides featuring textured fabric, metallic equestrian buckle, and cushioned insoles.",
					Tagline:     "Effortless comfort meets modern luxury.",
					Badge:       "Popular",
				},
				{
					ID:          "5",
					Name:        "Monogram Cross-Strap Leather Slides",
					Category:    "Footwear",
					Price:       "₦28,000",
					ImageURL:    "/static/images/quality.jpeg",
					Description: "Premium black leather cross-strap sliders adorned with signature gold-tone hardware branding.",
					Tagline:     "Bold design crafted for supreme comfort.",
					Badge:       "Featured",
				},
				{
					ID:          "6",
					Name:        "Vibrant Two-Tone Color Block Slides",
					Category:    "Footwear",
					Price:       "₦26,000",
					ImageURL:    "/static/images/foot-wears2.jpeg",
					Description: "Striking color-blocked leather slides combining rich blue uppers with vibrant orange cushioned footbeds.",
					Tagline:     "Stand out with vibrant, youthful design.",
					Badge:       "Limited Edition",
				},
				{
					ID:          "7",
					Name:        "Classic Black & Brown Comfort Set",
					Category:    "Footwear Collection",
					Price:       "₦27,000",
					ImageURL:    "/static/images/standardwear.jpeg",
					Description: "Handcrafted everyday leather slides featuring dual buckle straps and ergonomic cross-strap designs.",
					Tagline:     "Everyday durability with refined aesthetics.",
					Badge:       "Essential",
				},
				{
					ID:          "8",
					Name:        "Multi-Color Suede Block Slides",
					Category:    "Footwear",
					Price:       "₦27,000",
					ImageURL:    "/static/images/foots-wears1-3.jpeg",
					Description: "Artisanal multi-strip suede slides featuring bold purple, tan, orange, and deep green color bars.",
					Tagline:     "A masterpiece of color and texture.",
					Badge:       "Artisan Pick",
				},
				{
					ID:          "9",
					Name:        "Colorful Striped Designer Slides",
					Category:    "Footwear",
					Price:       "₦26,000",
					ImageURL:    "/static/images/foot-wears11-3.jpeg",
					Description: "Handcrafted comfort slides featuring layered textured fabric bands and cushioned insoles.",
					Tagline:     "Casual elegance for everyday wear.",
					Badge:       "Must-Have",
				},
			},
		}

		var page bytes.Buffer
		if err := tmpl.Execute(&page, data); err != nil {
			log.Printf("homepage template execution failed: %v", err)
			http.Error(w, "Unable to display this page right now.", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusOK)

		if r.Method == http.MethodGet {
			if _, err := page.WriteTo(w); err != nil {
				log.Printf("homepage response write failed: %v", err)
			}
		}
	}), nil
}
