package gemeni

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"github.com/joho/godotenv"
	"github.com/medali/ramo-server/internal/app"
	"google.golang.org/api/option"
)



func GemeniHelperFunc(infos app.HelperInput) interface{}{

	 err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }
	
	genApiKey := os.Getenv("GEN_API_KEY")




	us := "DigitalSpeak is your comprehensive digital marketing agency, offering all-encompassing solutions for online success. From website creation to media content management, copywriting, and media design, we provide everything your online business needs to thrive. Our team specializes in crafting engaging websites that stand out and function flawlessly, ensuring an exceptional user experience. We excel in creating captivating social media content, including posts, reels, and stories, to help you connect with your audience effectively. Additionally, our expert copywriters deliver compelling messaging that resonates with your target audience across various channels. With our media design services, your visual content will be both stunning and cohesive, reflecting your brand identity. At DigitalSpeak, we're dedicated to empowering your brand and helping you achieve your online goals with innovative strategies and unparalleled expertise."
	prompt := us + "\n\n  We've received the following details from \n"
	prompt += "- Business Name: " + infos.BusinessName + "\n" 
	prompt += "- Email: " + infos.Email + "\n"
    prompt += "- Industry: " + infos.Industry + "\n"
    prompt += "- Products/Services: " + infos.ProductsServices + "\n"
	if infos.HasWebsite {
        prompt += "- Website: " + infos.Website + "\n"
    } else {
        prompt += "The business  does not have a website.\n"
    }
	prompt += "- Active Social Media Platforms: " + infos.ActiveSocialMediaPlatforms + "\n"
	if infos.HasExistingBrandIdentity {
        prompt += ""
    } else {
        prompt += "The business  does not have an existing brand identity.\n"
    }
	prompt += "- Services of Interest: " + infos.ServicesOfInterest + "\n"
	prompt += "- Additional Information: " + infos.AdditionalInformation + "\n\n"

	   prompt += "Based on the information provided, Please generate text telling the business: " + infos.BusinessName + " How our agency DigitalSpeak would help enhance the business online presence and drive business growth and how can DigitalSpeak offer a range of services to elevate the brand and make " + infos.BusinessName + " business thrive," + "and provide the answer in " + infos.SpokenLanguage


	fmt.Println("#######################################################")
	fmt.Println(prompt)
	fmt.Println("#######################################################")

    ctx := context.Background()

	client, err := genai.NewClient(ctx, option.WithAPIKey(genApiKey))
    if err != nil {
    log.Fatal(err)
    }
    defer client.Close()

	model := client.GenerativeModel("gemini-pro")
	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
    if err != nil {
            log.Fatal(err)
    }

	response := resp.Candidates[0].Content
	return response
	
}