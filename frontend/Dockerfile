# Use a lightweight web server image
FROM nginx:alpine

# Remove default Nginx page
RUN rm -rf /usr/share/nginx/html/*

# Copy your static files into the web server directory
COPY template/ /usr/share/nginx/html/
# Optionally include testingUI if needed
# COPY testingUI/ /usr/share/nginx/html/testingUI/

# Expose port 80
EXPOSE 80

# Start Nginx
CMD ["nginx", "-g", "daemon off;"]