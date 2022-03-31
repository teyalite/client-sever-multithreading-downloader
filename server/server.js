import { fileTypeFromFile } from 'file-type';
import http from "http";
import { stat, createReadStream } from "fs";
import { pipeline } from 'stream';
import { promisify } from 'util';

const port = process.env.PORT || 3000;
const fileStat = promisify(stat);

http
    .createServer(async (req, res) => {
        try {

            const filename = req.url.slice(1);
            const { size } = await fileStat(filename);
            const type = await fileTypeFromFile(filename);

            if (req.method === 'HEAD') {
                /* Head request */
                res.writeHead(200, {
                    "Content-Length": size - 1,
                    "Content-Type": type.mime,
                    "Accept-Ranges": "bytes"
                });
                return res.end();
            }

            const range = req.headers.range;

            if (req.method === "GET" && range) {
                /* Get request */

                /* Extract range values from header */
                let [start, end] = range.replace(/bytes=/, "").split("-");

                start = parseInt(start, 10);
                end = end ? parseInt(end, 10) : size - 1;

                if (!isNaN(start) && isNaN(end)) {
                    /* if the start is a number but end is not a number */
                    start = start;
                    end = size - 1;
                }

                if (isNaN(start) && !isNaN(end)) {
                    /* if the end is a number but start is not a number */
                    start = size - end;
                    end = size - 1;
                }

                // unavailable range request
                if (start >= size || end >= size) {
                    // Return the 416 Range Not Satisfiable.
                    res.writeHead(416, {
                        "Content-Range": `bytes */${size}`
                    });
                    return res.end();
                }
                
                /** Send partial Content With HTTP Code 206 */
                res.writeHead(206, {
                    "Content-Range": `bytes ${start}-${end}/${size}`,
                    "Accept-Ranges": "bytes",
                    "Content-Length": end - start + 1,
                    "Content-Type": type
                });

                let readable = createReadStream(filename, { start: start, end: end });
                pipeline(readable, res, err => {});

                return;
            }

            res.writeHead(501);
            res.end('Not Implemented');

        } catch (error) {
            res.writeHead(404);
            res.end('NOT FOUND');
        }
    })
    .listen(port, () => console.log("Running on 3000 port"));