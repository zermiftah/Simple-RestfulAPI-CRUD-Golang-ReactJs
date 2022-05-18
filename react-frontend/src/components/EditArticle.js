import React, { useState, useEffect, useRef } from "react";
import { useParams } from 'react-router-dom'
import axios from "axios";

const EditArticle = () => {
    const [title, setTitle] = useState('');
    const [content, setContent] = useState('');
    const [category, setCategory] = useState('');
    const { id } = useParams();
    const [article, setArticle] = useState([])

    useEffect(() => {
        const fetcharticle = async () => {
            const data = await axios.get(`http://localhost:9999/article/${id}`)
            let temp = data.data.data;
            setArticle(temp);

        }
        fetcharticle()
    }, []);

    const handleUpdate = () => {
        axios.put(`http://localhost:9999/article/${id}`, {
            title: title,
            content: content,
            category: category
        })
    }

    const handleSubmit = (e) => {
        e.preventDefault();
        handleUpdate();
    }

    return (
        <>
            <div className="hidden sm:block" aria-hidden="true">
                <div className="py-5">
                    <div className="border-t border-gray-200" />
                </div>
            </div>

            <div className="mt-10 sm:mt-0">
                <div className="md:grid md:grid-cols-3 md:gap-6">
                    <div className="md:col-span-1">
                        <div className="px-4 sm:px-0">
                            <h3 className="text-lg font-medium leading-6 text-gray-900">Edit Article</h3>
                            <p className="mt-1 text-sm text-gray-600">Please edit this article.</p>
                        </div>
                    </div>
                    <div className="mt-5 md:mt-0 md:col-span-2">
                        <form action="#" method="POST" onSubmit={handleSubmit}>
                            <div className="shadow overflow-hidden sm:rounded-md">
                                <div className="px-4 py-5 bg-white sm:p-6">
                                    <div className="grid grid-cols-6 gap-6">
                                        <div className="col-span-6 sm:col-span-3">
                                            <label htmlFor="title" className="block text-sm font-medium text-gray-700">
                                                Title
                                            </label>
                                            <div className="mt-1">
                                                <input
                                                    id="title"
                                                    name="title"
                                                    rows={3}
                                                    className="shadow-sm focus:ring-indigo-500 focus:border-indigo-500 mt-1 block w-full sm:text-sm border border-gray-300 rounded-md"
                                                    placeholder={article.title}
                                                    onChange={(e) => { setTitle(e.target.value) }}
                                                />
                                            </div>
                                        </div>

                                        <div className="col-span-6">
                                            <label htmlFor="content" className="block text-sm font-medium text-gray-700">
                                                Content
                                            </label>
                                            <input
                                                type="text"
                                                placeholder={article.content}
                                                onChange={(e) => { setContent(e.target.value) }}
                                                name="content"
                                                id="content"
                                                autoComplete="street-address"
                                                className="shadow-sm focus:ring-indigo-500 focus:border-indigo-500 mt-1 block w-full sm:text-sm border border-gray-300 rounded-md"
                                            />
                                        </div>

                                        <div className="col-span-6 sm:col-span-3">
                                            <label htmlFor="category" className="block text-sm font-medium text-gray-700">
                                                Category
                                            </label>
                                            <div className="mt-1">
                                                <input
                                                    id="content"
                                                    name="content"
                                                    rows={3}
                                                    className="shadow-sm focus:ring-indigo-500 focus:border-indigo-500 mt-1 block w-full sm:text-sm border border-gray-300 rounded-md"
                                                    placeholder={article.category}
                                                    onChange={(e) => { setCategory(e.target.value) }}
                                                />
                                            </div>
                                        </div>
                                    </div>
                                </div>
                                <div className="px-4 py-3 bg-gray-50 text-right sm:px-6">
                                    <button
                                        name="submit" type="submit" onClick={handleUpdate}
                                        className="inline-flex justify-center py-2 px-4 border border-transparent shadow-sm text-sm font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
                                    >
                                        Save
                                    </button>
                                </div>
                            </div>
                        </form>
                    </div>
                </div>
            </div>

            <div className="hidden sm:block" aria-hidden="true">
                <div className="py-5">
                    <div className="border-t border-gray-200" />
                </div>
            </div>

        </>
    )
}

export default EditArticle;