use native_dialog::FileDialog;
use std::fs::File;
use std::io::copy;

use reqwest;

use crate::settings::SETTINGS;

pub fn download_and_save_file(url: &str) -> Result<(), Box<dyn std::error::Error>> {

    // Prompt the user to choose a save location
    let result = FileDialog::new()
        // .set_location(std::env::current_dir().unwrap().as_path())
        .set_location(&SETTINGS.images_folder)
        .add_filter("PNG Image", &["png"])
        .show_save_single_file()?; 

    // If user selected a path
    if let Some(file_path) = result {
        // Perform the file download and save operation

        // Send HTTP GET request to download the file
        let response = reqwest::blocking::get(url)?;
        // Open a file at the specified path for writing
        let mut file = File::create(file_path)?;

        // Copy the downloaded response body to the file
        copy(&mut response.bytes().unwrap().as_ref(), &mut file)?;

        // println!("File saved successfully to: {:?}", file_path);
    } else {
        // println!("User cancelled the dialog");
    }

    Ok(())
}
